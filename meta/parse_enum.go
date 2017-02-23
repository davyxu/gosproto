package meta

import (
	"errors"
)

func parseEnum(p *sprotoParser, fileD *FileDescriptor) {

	// enum
	p.Expect(Token_Enum)

	d := newDescriptor(fileD)

	// 名字
	d.Name = p.Expect(Token_Identifier).Value()

	// {
	p.Expect(Token_CurlyBraceL)

	for p.TokenID() != Token_CurlyBraceR {

		// 字段
		parseEnumField(p, d)

	}

	// }

	// 名字重复检查

	if fileD.NameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	fileD.addEnum(d)

}
