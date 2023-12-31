package test

import (
	"bytes"
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
	Expected string   `json:"expected"`
	Errors   []string `json:"errors"`
}

var scripts = "scripts"

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

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			suite := suite

			folder := filepath.Join(scripts, suite.Name)
			err = os.MkdirAll(folder, 0o700)
			if err != nil {
				t.Fatal(err)
			}

			for _, tc := range suite.Tests {
				if !only || tc.Only {
					t.Run(tc.Name, func(t *testing.T) {
						prog, errors, err := parser.Parse(string(tc.Code), suite.Name+"/"+tc.Name+".lsh", "./std", "../std")
						if err != nil {
							t.Fatal(err)
						}

						errors = append(errors, ast.Validate(prog)...)

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

						cmd := exec.Command("bash", file)

						// Create buffers to capture stdout and stderr
						stdoutBuf := &bytes.Buffer{}
						stderrBuf := &bytes.Buffer{}
						cmd.Stdout = stdoutBuf
						cmd.Stderr = stderrBuf

						// Run the command
						err = cmd.Run()
						if err != nil {
							t.Fatal(err)
						}

						// Capture stdout and stderr
						stdout := stdoutBuf.String()
						// stderr := stderrBuf.String()

						expect.Value(t, "result", stdout).ToBe(tc.Expected)
						// expect.Value(t, "errors", errs.Get()).ToCount(len(tc.Errors))
						// for i, erra := range errs.Get() {
						// 	if i < len(tc.Errors) {
						// 		expect.Value(t, "error", erra.Error()).ToBe(tc.Errors[i])
						// 	}
						// 	if i >= len(tc.Errors) {
						// 		t.Error(erra)
						// 	}
						// }
					})
				}
			}
		})
	}

	if only {
		t.Errorf("Test allways fails when only activated")
	}
}

var anum = regexp.MustCompile(`[^a-zA-Z0-9]`)

func asFile(s, ext string) string {
	return strings.ToLower(anum.ReplaceAllString(s, "_") + "." + ext)
}
