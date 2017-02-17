package meta

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	fileD, err := ParseFile("../example/addressbook.sp")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	fmt.Println(fileD.String())
}
