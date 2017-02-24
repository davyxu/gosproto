package meta

import "bytes"

type FileDescriptor struct {
	Structs []*Descriptor

	StructByName map[string]*Descriptor

	Enums []*Descriptor

	EnumByName map[string]*Descriptor

	ObjectsBySrcName map[string]*Descriptor
	Objects          []*Descriptor

	unknownFields []*lazyField
}

func (self *FileDescriptor) resolveAll() error {

	for _, v := range self.unknownFields {
		if _, err := v.resolve(2); err != nil {
			return err
		}
	}

	return nil
}

func (self *FileDescriptor) String() string {

	var bf bytes.Buffer

	bf.WriteString("Enum:")
	for _, st := range self.Enums {
		bf.WriteString(st.String())
		bf.WriteString("\n")
	}

	bf.WriteString("\n")

	bf.WriteString("Structs:")
	for _, st := range self.Structs {
		bf.WriteString(st.String())
		bf.WriteString("\n")
	}

	return bf.String()
}

func (self *FileDescriptor) NameExists(name string) bool {
	if _, ok := self.StructByName[name]; ok {
		return true
	}

	if _, ok := self.EnumByName[name]; ok {
		return true
	}

	return false
}

func (self *FileDescriptor) addObject(d *Descriptor, srcName string) {

	switch d.Type {
	case DescriptorType_Enum:
		self.Enums = append(self.Enums, d)
		self.EnumByName[d.Name] = d
	case DescriptorType_Struct:
		self.Structs = append(self.Structs, d)
		self.StructByName[d.Name] = d
	}

	d.SrcName = srcName

	self.Objects = append(self.Objects, d)

	self.ObjectsBySrcName[srcName] = d
}

func NewFileDescriptor() *FileDescriptor {

	return &FileDescriptor{
		StructByName:     make(map[string]*Descriptor),
		EnumByName:       make(map[string]*Descriptor),
		ObjectsBySrcName: make(map[string]*Descriptor),
	}

}
