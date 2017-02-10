package meta

import (
	"bytes"
)

type Descriptor struct {
	Name string

	Fields      []*FieldDescriptor
	FieldByName map[string]*FieldDescriptor

	File *FileDescriptor
}

func (self *Descriptor) String() string {

	var bf bytes.Buffer

	bf.WriteString(self.Name)

	bf.WriteString(":\n")

	for _, fd := range self.Fields {
		bf.WriteString("	")
		bf.WriteString(fd.String())
		bf.WriteString("\n")
	}

	bf.WriteString("\n")

	return bf.String()
}

func (self *Descriptor) AddField(fd *FieldDescriptor) {
	self.Fields = append(self.Fields, fd)
	self.FieldByName[fd.Name] = fd
}

func NewDescriptor(f *FileDescriptor) *Descriptor {
	return &Descriptor{
		File:        f,
		FieldByName: make(map[string]*FieldDescriptor),
	}
}
