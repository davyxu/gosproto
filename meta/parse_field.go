package meta

import "errors"

func parseField(p *sprotoParser, d *Descriptor) {

	fd := NewFieldDescriptor(d)

	// 字段名
	fd.Name = p.Expect(Token_Identifier).Value()

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New("Duplicate field name: " + d.Name))
	}

	// tag
	fd.Tag = p.Expect(Token_Numeral).ToInt()

	// :
	p.Expect(Token_Colon)

	var typeName string

	switch p.TokenID() {
	// 数组
	case Token_Star:
		p.NextToken()

		fd.Repeatd = true

		typeName = p.Expect(Token_Identifier).Value()

	case Token_Identifier:
		// 普通字段
		typeName = p.TokenValue()
		p.NextToken()
		break
	default:
	}

	// 根据类型名查找类型及结构体类型
	fd.Type, fd.Complex = fd.parseType(typeName)

	if fd.Type == FieldType_None {
		panic(errors.New("type not found: " + typeName))
	}

	// map的索引解析 (
	if fd.Complex != nil && p.TokenID() == Token_ParenL {
		p.NextToken()

		// 索引的字段
		mainIndexName := p.Expect(Token_Identifier).Value()

		if indexFd, ok := fd.Complex.FieldByName[mainIndexName]; ok {
			fd.MainIndex = indexFd
		} else {
			panic(errors.New("Main index not found:" + typeName))
		}

		p.Expect(Token_ParenR)

	}
	// )

	d.AddField(fd)
}
