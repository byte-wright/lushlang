package test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/akabio/expect"
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/bash"
	"github.com/byte-wright/lush/internal/bcode"
	"github.com/byte-wright/lush/internal/parser"
	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"gopkg.in/yaml.v2"
)

type TestSuite struct {
	Name  string  `json:"name"`
	Tests []*Test `json:"tests"`
}

type Test struct {
	Name     string   `json:"name"`
	Only     bool     `json:"only"`
	Code     string   `json:"code"`
	Args     []string `json:"args"`
	Expected string   `json:"expected"`
	Stderr   string   `json:"stderr"`
	Exit     int      `json:"exit"`
	Errors   []string `json:"errors"`
}

var scripts = "scripts"

// 3.2 is the oldest supported bash, macOS still ships it
var images = []string{"bash:3.2", "bash:5.3"}

var env = []string{"LUSH_TEST=hello", "lush_test=env"}

type runner func(t *testing.T, file string, args []string) (stdout, stderr string, exit int)

func TestAll(t *testing.T) {
	os.RemoveAll(scripts)

	files, err := os.ReadDir("./")
	if err != nil {
		t.Fatal(err)
	}
	suites := []*TestSuite{}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".yaml") {
			d, err := os.ReadFile(name)
			if err != nil {
				t.Fatal(err)
			}
			suite := &TestSuite{}
			err = yaml.Unmarshal(d, &suite)
			if err != nil {
				t.Fatal(err)
			}
			suite.Name = name[:len(name)-5]
			suites = append(suites, suite)
		}
	}

	only := false
	for _, suite := range suites {
		for _, t := range suite.Tests {
			if t.Only {
				only = true
			}
		}
	}

	t.Run("local", func(t *testing.T) {
		runSuites(t, suites, only, runLocal)
	})

	for _, image := range images {
		t.Run(image, func(t *testing.T) {
			testcontainers.SkipIfProviderIsNotHealthy(t)
			runSuites(t, suites, only, startContainer(t, image))
		})
	}

	if only {
		t.Errorf("Test allways fails when only activated")
	}
}

func runSuites(t *testing.T, suites []*TestSuite, only bool, run runner) {
	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			folder := filepath.Join(scripts, suite.Name)
			err := os.MkdirAll(folder, 0o700)
			if err != nil {
				t.Fatal(err)
			}

			for _, tc := range suite.Tests {
				if !only || tc.Only {
					t.Run(tc.Name, func(t *testing.T) {
						t.Parallel()

						prog, errors, err := parser.Parse(string(tc.Code), suite.Name+"/"+tc.Name+".lsh", "./std", os.DirFS("../std"))
						if err != nil {
							t.Fatal(err)
						}

						expectedErrors := map[string]bool{}
						for _, e := range tc.Errors {
							expectedErrors[e] = true
						}

						for _, err := range errors {
							if !expectedErrors[err.Error()] {
								t.Errorf("unexpected error `%v`", err.Error())
							}
							delete(expectedErrors, err.Error())
						}

						for expe := range expectedErrors {
							t.Errorf("expected error `%v`", expe)
						}

						if len(errors) > 0 {
							return
						}

						err = os.WriteFile(filepath.Join(folder, asFile(tc.Name, "ast")), []byte(ast.Print(prog)), 0o600)
						if err != nil {
							t.Fatal(err)
						}

						bc, err := bcode.Compile(prog)
						if err != nil {
							t.Fatal(err)
						}

						err = os.WriteFile(filepath.Join(folder, asFile(tc.Name, "basm")), []byte(bcode.Print(bc)), 0o600)
						if err != nil {
							t.Fatal(err)
						}

						file := filepath.Join(folder, asFile(tc.Name, "sh"))
						err = os.WriteFile(file, []byte(bash.Translate(bc)), 0o700)
						if err != nil {
							t.Fatal(err)
						}

						stdout, stderr, exit := run(t, file, tc.Args)

						expect.Value(t, "result", stdout).ToBe(tc.Expected)
						expect.Value(t, "stderr", stderr).ToBe(tc.Stderr)
						expect.Value(t, "exit code", exit).ToBe(tc.Exit)
					})
				}
			}
		})
	}
}

func runLocal(t *testing.T, file string, args []string) (string, string, int) {
	cmd := exec.Command("bash", append([]string{file}, args...)...)
	cmd.Env = append(os.Environ(), env...)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	exit := 0
	if exitErr, isExitErr := err.(*exec.ExitError); isExitErr {
		exit = exitErr.ExitCode()
	} else if err != nil {
		t.Fatal(err)
	}

	return stdout.String(), stderr.String(), exit
}

func startContainer(t *testing.T, image string) runner {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	ctr, err := testcontainers.Run(context.Background(), image,
		testcontainers.WithCmd("sleep", "infinity"),
		testcontainers.WithHostConfigModifier(func(hc *container.HostConfig) {
			// a mount instead of copies: the scripts are written while the tests run,
			// and some tests read files from the tests folder
			hc.Binds = []string{wd + ":/tests:ro"}
		}),
	)
	// sleep ignores SIGTERM as the container's first process, stopping would wait for 10s
	testcontainers.CleanupContainer(t, ctr, testcontainers.StopTimeout(0))
	if err != nil {
		t.Fatal(err)
	}

	return func(t *testing.T, file string, args []string) (string, string, int) {
		exit, output, err := ctr.Exec(context.Background(), append([]string{"bash", file}, args...),
			tcexec.WithWorkingDir("/tests"), tcexec.WithEnv(env))
		if err != nil {
			t.Fatal(err)
		}

		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		_, err = stdcopy.StdCopy(stdout, stderr, output)
		if err != nil {
			t.Fatal(err)
		}

		return stdout.String(), stderr.String(), exit
	}
}

var anum = regexp.MustCompile(`[^a-zA-Z0-9]`)

func asFile(s, ext string) string {
	return strings.ToLower(anum.ReplaceAllString(s, "_") + "." + ext)
}
