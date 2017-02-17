package meta

import (
	"fmt"
	"io/ioutil"

	"github.com/davyxu/golexer"
)

// 从文件解析
func ParseFile(fileName string) (fileD *FileDescriptor, retErr error) {

	var data []byte

	data, retErr = ioutil.ReadFile(fileName)
	if retErr != nil {
		return
	}

	fileD, retErr = ParseString(string(data))
	if retErr != nil {
		fmt.Printf("parse %s failed\n", fileName)
		return
	}

	fileD.FileName = fileName

	return
}

// 解析字符串
func ParseString(data string) (fileD *FileDescriptor, retErr error) {

	p := newSProtoParser()

	defer golexer.ErrorCatcher(func(err error) {

		line, _ := p.TokenPos()

		fmt.Printf("line %d \n", line)

		retErr = err

	})

	fileD = NewFileDescriptor()

	p.Lexer().Start(data)

	p.NextToken()

	for p.TokenID() != Token_EOF {

		parseStruct(p, fileD)

		p.NextToken()

	}

	p.resolveAll()

	return fileD, nil
}
