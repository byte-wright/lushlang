package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/byte-wright/lush/internal/bash"
	"github.com/byte-wright/lush/internal/bcode"
)

func TestExample(t *testing.T) {
	examples := "../../examples"

	fs, err := os.ReadDir(examples)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range fs {
		if !strings.HasSuffix(f.Name(), ".lsh") {
			continue
		}

		fmt.Println(f)
		dat, err := os.ReadFile(filepath.Join(examples, f.Name()))
		if err != nil {
			t.Error(err)
		}

		prog, err := Parse(string(dat), f.Name(), "./std")
		if err != nil {
			t.Error(err)
		}

		for _, imp := range prog.Libs {
			fmt.Println("import", imp.Path)
		}

		// fmt.Println(prog.Print())

		// yd, err := yaml.Marshal(prog)
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// fmt.Println(string(yd))

		bc, err := bcode.Compile(prog)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(bc.Print())

		bs := bash.Translate(bc)
		// fmt.Println(bs)

		nName := strings.TrimSuffix(f.Name(), ".lsh") + ".sh"

		err = os.WriteFile(filepath.Join(examples, nName), []byte(bs), 0o700)
		if err != nil {
			t.Fatal(err)
		}

	}

	t.Fail()
}
