package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davyxu/gosproto/meta"
)

var paramOut = flag.String("out", "", "output filename")
var paramGoPackage = flag.String("gopackage", "", "package name in go files")
var paramType = flag.String("type", "", "output file type")

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
		gen_go(fileD, *paramGoPackage, *paramOut)
	case "sproto":
		gen_sproto(fileD, *paramOut)
	default:
		fmt.Println("unknown out file type: ", *paramType)
		os.Exit(1)
	}

}
