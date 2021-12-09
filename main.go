package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/profsergiocosta/ccompiler/gen"
	"github.com/profsergiocosta/ccompiler/parser"
)

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		abs, _ := filepath.Abs(path)
		fmt.Printf("Could not find file or directory: %s \n", abs)
		os.Exit(1)
	}
	return fileInfo.IsDir()
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func main() {

	arg := os.Args[1:]

	if len(arg) == 1 {
		path := arg[0]

		p := parser.New(path)
		program := p.ParseProgram()

		fmt.Println(program)
		s := gen.Generate(program)

		outFileName := FilenameWithoutExtension(path)

		out, _ := os.Create(outFileName + ".s")

		out.WriteString(s)

	}

}
