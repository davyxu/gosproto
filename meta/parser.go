package meta

import (
	"fmt"
	"io/ioutil"

	"github.com/davyxu/golexer"
)

func ParseFile(fileName string) error {
	fileD := NewFileDescriptor()

	err := rawPaseFile(fileD, fileName)
	if err != nil {
		return err
	}

	return fileD.resolveAll()
}

func ParseFileList(fileD *FileDescriptor, filelist []string) (string, error) {

	for _, filename := range filelist {
		if err := rawPaseFile(fileD, filename); err != nil {
			return filename, err
		}
	}

	return "", fileD.resolveAll()

}

func ParseString(data string) error {

	fileD := NewFileDescriptor()

	if err := rawParse(fileD, data); err != nil {
		return err
	}

	return fileD.resolveAll()
}

// 从文件解析
func rawPaseFile(fileD *FileDescriptor, fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return rawParse(fileD, string(data))
}

// 解析字符串
func rawParse(fileD *FileDescriptor, data string) (retErr error) {

	p := newSProtoParser()

	defer golexer.ErrorCatcher(func(err error) {

		line, _ := p.TokenPos()

		fmt.Printf("line %d \n", line)

		retErr = err

	})

	p.Lexer().Start(data)

	p.NextToken()

	for p.TokenID() != Token_EOF {

		parseStruct(p, fileD)

		p.NextToken()

	}

	return nil
}
