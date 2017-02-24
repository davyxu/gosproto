package meta

import "fmt"

type FieldDescriptor struct {
	*CommentGroup
	Name      string
	Type      FieldType
	Tag       int
	AutoTag   int
	Repeatd   bool
	MainIndex *FieldDescriptor
	Complex   *Descriptor

	Struct *Descriptor
}

func (self *FieldDescriptor) TagNumber() int {

	if self.AutoTag == -1 {
		return self.Tag
	}

	return self.AutoTag
}

func (self *FieldDescriptor) TypeString() string {
	return self.typeStr(false)
}

func (self *FieldDescriptor) CompatibleTypeString() string {
	return self.typeStr(true)
}

func (self *FieldDescriptor) typeStr(compatible bool) (ret string) {

	if self.Repeatd {
		ret = "*"
	}

	if compatible {
		ret += self.CompatibleTypeName()
	} else {
		ret += self.TypeName()
	}

	if self.MainIndex != nil {
		ret += fmt.Sprintf("(%s)", self.MainIndex.Name)
	}

	return
}

func (self *FieldDescriptor) String() string {

	return fmt.Sprintf("%s %d : %s", self.Name, self.TagNumber(), self.TypeString())
}

func (self *FieldDescriptor) Kind() string {

	return self.Type.String()
}

func (self *FieldDescriptor) CompatibleTypeName() string {

	switch self.Type {

	case FieldType_Struct:
		return self.Complex.Name
	case FieldType_Int32,
		FieldType_Int64,
		FieldType_UInt32,
		FieldType_UInt64,
		FieldType_Enum:
		return FieldType_Integer.String()
	default:
		return self.Type.String()
	}

}

func (self *FieldDescriptor) TypeName() string {

	switch self.Type {

	case FieldType_Struct, FieldType_Enum:
		return self.Complex.Name
	default:
		return self.Type.String()
	}

}

func (self *FieldDescriptor) parseType(name string) (ft FieldType, structType *Descriptor) {

	ft = ParseFieldType(name)

	if ft != FieldType_None {
		return ft, nil
	}

	if d, ok := self.Struct.File.StructByName[name]; ok {
		return FieldType_Struct, d
	}

	if d, ok := self.Struct.File.EnumByName[name]; ok {
		return FieldType_Enum, d
	}

	return FieldType_None, nil

}

func newFieldDescriptor(d *Descriptor) *FieldDescriptor {
	return &FieldDescriptor{
		CommentGroup: newCommentGroup(),
		Struct:       d,
		AutoTag:      -1,
	}
}
