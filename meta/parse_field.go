package meta

import "errors"

type parsingField struct {
	typeName      string
	mainIndexName string

	fd *FieldDescriptor
}

func (self *parsingField) resolve(pass int) bool {

	if self.fd.Type == FieldType_None {
		if pass > 1 {

			panic(errors.New("type not found: " + self.typeName))
		} else {
			return false
		}
	}

	if self.mainIndexName != "" {
		if indexFd, ok := self.fd.Complex.FieldByName[self.mainIndexName]; ok {
			self.fd.MainIndex = indexFd
		} else {
			if pass > 1 {
				panic(errors.New("Main index not found:" + self.typeName))
			} else {
				return false
			}

		}
	}

	return true
}

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

	pf := &parsingField{typeName: typeName, fd: fd}

	// map的索引解析 (
	if p.TokenID() == Token_ParenL {
		p.NextToken()

		// 索引的字段
		pf.mainIndexName = p.Expect(Token_Identifier).Value()

		p.Expect(Token_ParenR)

	}
	// )

	// 尝试首次解析
	if !pf.resolve(1) {
		p.unknownFields = append(p.unknownFields)
	}

	d.AddField(fd)
}
