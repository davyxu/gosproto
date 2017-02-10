package meta

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	fileD, err := ParseFile("../exsample/test.sproto")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	fmt.Println(fileD.String())
}
