package meta

import "errors"

func parseEnumField(p *sprotoParser, d *Descriptor) {

	fd := newFieldDescriptor(d)

	nameToken := p.RawToken()

	// 字段名
	fd.Name = p.Expect(Token_Identifier).Value()

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New("Duplicate field name: " + d.Name))
	}

	// =
	p.Expect(Token_Assign)

	// tag
	fd.Tag = p.Expect(Token_Numeral).ToInt()

	fd.Type = FieldType_Int32

	fd.CommentGroup = p.CommentGroupByLine(nameToken.Line())

	checkField(d, fd)

	d.addField(fd)

	return
}
