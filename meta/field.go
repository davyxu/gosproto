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

func (self *FieldDescriptor) String() string {

	var starStr string

	if self.Repeatd {
		starStr = "*"
	}

	var indexStr string
	if self.MainIndex != nil {
		indexStr = fmt.Sprintf("(%s)", self.MainIndex.Name)
	}

	return fmt.Sprintf("%s %d : %s%s%s", self.Name, self.Tag, starStr, self.TypeName(), indexStr)
}

func (self *FieldDescriptor) Kind() string {

	return self.Type.String()
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
