package meta

import "fmt"

type FieldDescriptor struct {
	Name      string
	Type      FieldType
	Tag       int
	Repeatd   bool
	MainIndex *FieldDescriptor
	Complex   *Descriptor

	Struct *Descriptor
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

	return fmt.Sprintf("%s %d : %s", self.Name, self.Tag, self.TypeString())
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
		FieldType_UInt64:
		return FieldType_Integer.String()
	default:
		return self.Type.String()
	}

}

func (self *FieldDescriptor) TypeName() string {

	switch self.Type {

	case FieldType_Struct:
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

	return FieldType_None, nil

}

func NewFieldDescriptor(d *Descriptor) *FieldDescriptor {
	return &FieldDescriptor{
		Struct: d,
	}
}
