package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExample(t *testing.T) {
	examples := "../../examples"

	fs, err := os.ReadDir(examples)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range fs {
		fmt.Println(f)
		dat, err := os.ReadFile(filepath.Join(examples, f.Name()))
		if err != nil {
			t.Error(err)
		}

		prog, err := Parse(string(dat), f.Name())
		if err != nil {
			t.Error(err)
		}

		fmt.Println(prog.Print())

	}

	t.Fail()
}
