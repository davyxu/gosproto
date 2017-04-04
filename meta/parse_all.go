package meta

import (
	"errors"
	"fmt"
	"github.com/davyxu/golexer"
	"io/ioutil"
)

func ParseFile(fileName string) (*FileDescriptor, error) {
	fileD := NewFileDescriptor()

	err := rawPaseFile(fileD, fileName)
	if err != nil {
		return nil, err
	}

	return fileD, fileD.resolveAll()
}

func ParseFileList(fileD *FileDescriptor, filelist []string) (string, error) {

	for _, filename := range filelist {
		if err := rawPaseFile(fileD, filename); err != nil {
			return filename, err
		}
	}

	return "", fileD.resolveAll()

}

func ParseString(data string) (*FileDescriptor, error) {

	fileD := NewFileDescriptor()

	if err := rawParse(fileD, data, data); err != nil {
		return nil, err
	}

	return fileD, fileD.resolveAll()
}

// 从文件解析
func rawPaseFile(fileD *FileDescriptor, fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return rawParse(fileD, string(data), fileName)
}

// 解析字符串
func rawParse(fileD *FileDescriptor, data string, srcName string) (retErr error) {

	p := newSProtoParser(srcName)

	defer golexer.ErrorCatcher(func(err error) {

		retErr = fmt.Errorf("%s %s", p.PreTokenPos().String(), err.Error())

	})

	p.Lexer().Start(data)

	p.NextToken()

	for p.TokenID() != Token_EOF {

		switch p.TokenID() {
		case Token_Dot, Token_Message:
			parseStruct(p, fileD, srcName)
		case Token_Enum:
			parseEnum(p, fileD, srcName)
		default:
			panic(errors.New("Expect '.' or 'enum'"))
		}

		p.NextToken()

	}

	return nil
}
