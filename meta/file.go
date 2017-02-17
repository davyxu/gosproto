package meta

import "bytes"

type FileDescriptor struct {
	Structs []*Descriptor

	StructByName map[string]*Descriptor

	unknownFields []*parsingField
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

	for _, st := range self.Structs {
		bf.WriteString(st.String())
		bf.WriteString("\n")
	}

	bf.WriteString("\n")

	return bf.String()
}

func (self *FileDescriptor) AddStruct(d *Descriptor) {
	self.Structs = append(self.Structs, d)
	self.StructByName[d.Name] = d
}
func NewFileDescriptor() *FileDescriptor {

	return &FileDescriptor{
		StructByName: make(map[string]*Descriptor),
	}

}
