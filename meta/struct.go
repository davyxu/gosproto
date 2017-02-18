package meta

import (
	"bytes"
	"errors"
)

type Descriptor struct {
	Name string

	Fields      []*FieldDescriptor
	FieldByName map[string]*FieldDescriptor
	FieldByTag  map[int]*FieldDescriptor

	File *FileDescriptor
}

// c# 要使用的fieldcount
func (self *Descriptor) MaxFieldCount() int {
	maxn := len(self.Fields)
	lastTag := -1

	for _, fd := range self.Fields {
		if fd.Tag < lastTag {
			panic(errors.New("tag must in ascending order"))
		}

		if fd.Tag > lastTag+1 {
			maxn++
		}

		lastTag = fd.Tag
	}

	return maxn
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

func (self *Descriptor) addField(fd *FieldDescriptor) {
	self.Fields = append(self.Fields, fd)
	self.FieldByName[fd.Name] = fd
	self.FieldByTag[fd.Tag] = fd
}

func newDescriptor(f *FileDescriptor) *Descriptor {
	return &Descriptor{
		File:        f,
		FieldByName: make(map[string]*FieldDescriptor),
		FieldByTag:  make(map[int]*FieldDescriptor),
	}
}
