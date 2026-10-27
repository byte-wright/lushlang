package lush

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/akabio/expect"
)

func TestToBashDebugOutsideRepo(t *testing.T) {
	root := t.TempDir()

	err := os.Mkdir(filepath.Join(root, "greet"), 0o700)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(root, "greet", "greet.lsh"), []byte(`
func hello(name string) string {
  return "hello " + name
}
`), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// neither root nor std may depend on the working directory
	err = os.Chdir(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)

	_, _, sh, err := ToBashDebug(`
import "strings"
import "/greet"

println(greet.hello("world"), strings.hasPrefix("world", "w"))
`, "main.lsh", root)
	if err != nil {
		t.Fatal(err)
	}

	out, err := exec.Command("bash", "-c", sh).Output()
	if err != nil {
		t.Fatal(err)
	}

	expect.Value(t, "output", string(out)).ToBe("hello world true\n")
}
