package meta

import (
	"errors"
)

func parseStruct(p *sprotoParser, fileD *FileDescriptor) {

	// .
	p.Expect(Token_Dot)

	d := newDescriptor(fileD)

	// 名字
	d.Name = p.Expect(Token_Identifier).Value()

	// {
	p.Expect(Token_CurlyBraceL)

	for p.TokenID() != Token_CurlyBraceR {

		// 字段
		parseField(p, d)

	}

	// }

	// 结构体重复检查
	if _, ok := fileD.StructByName[d.Name]; ok {
		panic(errors.New("Duplicate struct name: " + d.Name))
	}

	fileD.AddStruct(d)

}
