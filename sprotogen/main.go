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

func mergeSchema(filelist []string) (ret []*meta.FileDescriptor) {

	if len(filelist) == 0 {
		fmt.Println("require sproto file")
		os.Exit(1)
	}

	for _, filename := range filelist {
		fileD, err := meta.ParseFile(filename)

		if err != nil {
			fmt.Println("parse failed, ", err.Error())
			os.Exit(1)
		}

		ret = append(ret, fileD)
	}

	return
}

func main() {

	flag.Parse()

	fileDList := mergeSchema(flag.Args())

	switch *paramType {
	case "go":
		gen_go(fileDList, *paramGoPackage, *paramOut)
	default:
		fmt.Println("unknown out file type: ", *paramType)
		os.Exit(1)
	}

}
