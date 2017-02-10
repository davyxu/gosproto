package meta

import (
	"fmt"
)

type FieldType int

const (
	FieldType_None = iota
	FieldType_Integer
	FieldType_Bool
	FieldType_String
	FieldType_Struct
)

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
	switch self.Type {
	case FieldType_Bool:
		return "bool"
	case FieldType_Integer:
		return "integer"
	case FieldType_String:
		return "string"
	case FieldType_Struct:
		return "struct"
	}

	return "!none!"
}

func (self *FieldDescriptor) TypeName() string {

	switch self.Type {
	case FieldType_Bool:
		return "bool"
	case FieldType_Integer:
		return "integer"
	case FieldType_String:
		return "string"
	case FieldType_Struct:
		return self.Complex.Name
	}

	return "!none!"
}

func (self *FieldDescriptor) parseType(name string) (ft FieldType, structType *Descriptor) {

	switch name {
	case "integer":
		return FieldType_Integer, nil
	case "string":
		return FieldType_String, nil
	case "bool":
		return FieldType_Bool, nil
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
