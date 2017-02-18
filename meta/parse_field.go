package meta

import "errors"

type parsingField struct {
	typeName      string
	mainIndexName string

	fd *FieldDescriptor

	miss bool
}

func (self *parsingField) resolve(pass int) (bool, error) {

	self.fd.Type, self.fd.Complex = self.fd.parseType(self.typeName)

	if self.fd.Type == FieldType_None {
		if pass > 1 {

			return true, errors.New("type not found: " + self.typeName)
		} else {

			self.miss = true
			return true, nil
		}
	}

	if self.mainIndexName != "" {
		if indexFd, ok := self.fd.Complex.FieldByName[self.mainIndexName]; ok {
			self.fd.MainIndex = indexFd
		} else {
			if pass > 1 {
				return true, errors.New("Main index not found:" + self.typeName)
			} else {
				return true, nil
			}

		}
	}

	return false, nil
}

type fieldParseType int

const (
	fieldParseType_None fieldParseType = iota
	fieldParseType_StructField
	fieldParseType_EnumField
)

func parseField(p *sprotoParser, d *Descriptor) (fpt fieldParseType) {

	fd := newFieldDescriptor(d)

	// 字段名
	fd.Name = p.Expect(Token_Identifier).Value()

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New("Duplicate field name: " + d.Name))
	}

	// tag
	fd.Tag = p.Expect(Token_Numeral).ToInt()

	// :
	if p.TokenID() == Token_Colon {

		p.NextToken()

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
		if need2Pass, _ := pf.resolve(1); need2Pass {
			d.File.unknownFields = append(d.File.unknownFields, pf)
		}

		fpt = fieldParseType_StructField
	} else {

		fd.Type = FieldType_Int32

		fpt = fieldParseType_EnumField

	}

	d.AddField(fd)

	return
}
