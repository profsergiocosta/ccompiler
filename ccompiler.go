package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/profsergiocosta/ccompiler/eval"
	"github.com/profsergiocosta/ccompiler/parser"
)

// https://golang.org/doc/code.html

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

		if IsDirectory(path) {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}
			for _, f := range files {

				if filepath.Ext(f.Name()) == ".jack" {
					abs, _ := filepath.Abs(path + "/" + f.Name())
					fmt.Printf("Compiling: %s \n", abs)
					p := parser.New(abs)
					p.Compile()
				}

			}

		} else {
			//abs, _ := filepath.Abs(path)

			//dirname, _ := filepath.Split(abs)
			
			//fmt.Printf("Compiling: %s  \n", abs)
			p := parser.New(path)
			program := p.ParseProgram()
			//fmt.Printf(program.Function.TokenLiteral())
			//fmt.Printf(program.Function.Statement.String())
			s := eval.Eval(program)
			//out := *os.File
			outFileName := FilenameWithoutExtension(path)
			
			out, _ := os.Create(outFileName+".s")
			//fmt.Printf(s)
			out.WriteString(s)
			//out.Close();
			//p.Compile()
			//p.CompileExpression()
			//p.CompileIf()

		}

	}

}
