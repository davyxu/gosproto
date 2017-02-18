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

	var lastFpt fieldParseType

	for p.TokenID() != Token_CurlyBraceR {

		// 字段
		fpt := parseField(p, d)
		if lastFpt != fieldParseType_None && lastFpt != fpt {
			panic(errors.New("invalid enum field, keep no colon!"))
		}

		lastFpt = fpt

	}

	// }

	// 名字重复检查

	if fileD.NameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	switch lastFpt {
	case fieldParseType_EnumField:
		fileD.addEnum(d)
	case fieldParseType_StructField:
		fileD.addStruct(d)
	}

}
