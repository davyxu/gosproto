package main

import (
	"fmt"
	"github.com/davyxu/gosproto/meta"
	"os"
	"strconv"
	"strings"
)

// 本文件内功能仅做项目内部功能使用, 不做通用功能

func enumValueOffset(fileset *meta.FileDescriptorSet) {

	var allTagNumbers = map[int]*meta.FieldDescriptor{}

	for _, file := range fileset.Files {

		for _, e := range file.Enums {

			if rawValue, ok := e.MatchTag("EnumValueOffset"); ok {

				if offset, err := strconv.ParseInt(strings.TrimSpace(rawValue), 10, 32); err == nil {

					e.TagBase = int(offset)
				}

			}

			if strings.HasSuffix(e.Name, "Result") {

				for _, fd := range e.Fields {
					if fd.TagNumber() == 0 {
						continue
					}

					if prev, ok := allTagNumbers[fd.TagNumber()]; ok {

						fmt.Printf("Duplicated enum value: %d  %s.%s(%s)  prev: %s.%s(%s)\n", fd.TagNumber(), fd.Struct.Name, fd.Name, fd.Struct.File.FileName, prev.Struct.Name, prev.Name, prev.Struct.File.FileName)
						os.Exit(1)
					}

					allTagNumbers[fd.TagNumber()] = fd
				}

			}

		}
	}

}
