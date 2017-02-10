package example

import (
	"reflect"
	"testing"

	"github.com/davyxu/gosproto"
)

var abData []byte = []byte{
	1, 0, 0, 0, 122, 0, 0, 0,
	68, 0, 0, 0, 4, 0, 0,
	0, 34, 78, 1, 0, 0, 0,
	5, 0, 0, 0, 65, 108, 105,
	99, 101, 45, 0, 0, 0, 19,
	0, 0, 0, 2, 0, 0, 0,
	4, 0, 9, 0, 0, 0, 49,
	50, 51, 52, 53, 54, 55, 56,
	57, 18, 0, 0, 0, 2, 0,
	0, 0, 6, 0, 8, 0, 0,
	0, 56, 55, 54, 53, 52, 51,
	50, 49, 46, 0, 0, 0, 4,
	0, 0, 0, 66, 156, 1, 0,
	0, 0, 3, 0, 0, 0, 66,
	111, 98, 25, 0, 0, 0, 21,
	0, 0, 0, 2, 0, 0, 0,
	8, 0, 11, 0, 0, 0, 48,
	49, 50, 51, 52, 53, 54, 55,
	56, 57, 48,
}

func TestAddressBook(t *testing.T) {

	for _, tp := range SProtoStructs {
		t.Log(tp.Name())
	}

	input := &AddressBook{
		Person: []*Person{
			&Person{
				Name: "Alice",
				Id:   10000,
				Phone: []*PhoneNumber{
					&PhoneNumber{
						Number: "123456789",
						Type:   1,
					},
					&PhoneNumber{
						Number: "87654321",
						Type:   2,
					},
				},
			},
			&Person{
				Name: "Bob",
				Id:   20000,
				Phone: []*PhoneNumber{
					&PhoneNumber{
						Number: "01234567890",
						Type:   3,
					},
				},
			},
		},
	}

	data, err := sproto.Encode(input)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var sample AddressBook
	_, err = sproto.Decode(data, &sample)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(abData, data) || !reflect.DeepEqual(&sample, input) {
		t.FailNow()
	}

}
