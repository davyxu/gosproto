package meta

import (
	"errors"
)

func parseStruct(p *sprotoParser, fileD *FileDescriptor) {

	// .
	dotToken := p.Expect(Token_Dot)

	d := newDescriptor(fileD)

	// 名字
	d.Name = p.Expect(Token_Identifier).Value()

	d.CommentGroup = p.CommentGroupByLine(dotToken.Line())

	// {
	p.Expect(Token_CurlyBraceL)

	for p.TokenID() != Token_CurlyBraceR {

		// 字段
		parseField(p, d)

	}

	// }

	// 名字重复检查

	if fileD.NameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	fileD.addStruct(d)

}
