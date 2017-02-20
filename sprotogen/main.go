package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davyxu/gosproto/meta"
)

var paramOut = flag.String("out", "", "output filename")
var paramPackage = flag.String("package", "", "package name in go files")
var paramType = flag.String("type", "", "output file type")
var paramCellnetReg = flag.Bool("cellnet_reg", false, "for type go, generate sproto auto register entry for github.com/davyxu/cellnet")

func mergeSchema(filelist []string) *meta.FileDescriptor {

	if len(filelist) == 0 {
		fmt.Println("require sproto file")
		os.Exit(1)
	}

	fileD := meta.NewFileDescriptor()
	errorFileName, err := meta.ParseFileList(fileD, filelist)
	if err != nil {
		fmt.Println(errorFileName, err.Error())
		os.Exit(1)
	}

	return fileD
}

func main() {

	flag.Parse()

	fileD := mergeSchema(flag.Args())

	switch *paramType {
	case "go":
		gen_go(fileD, *paramPackage, *paramOut, *paramCellnetReg)
	case "sproto":
		gen_sproto(fileD, *paramOut)
	case "cs":
		gen_csharp(fileD, *paramPackage, *paramOut)
	case "lua":
		gen_lua(fileD, *paramPackage, *paramOut)
	default:
		fmt.Println("unknown out file type: ", *paramType)
		os.Exit(1)
	}

}
