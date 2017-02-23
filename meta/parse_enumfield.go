package meta

import "errors"

func parseEnumField(p *sprotoParser, d *Descriptor) {

	fd := newFieldDescriptor(d)

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

	checkField(d, fd)

	d.addField(fd)

	return
}
