# gosproto
[sproto](https://github.com/cloudwu/sproto) golang implement

# Based on code

https://github.com/xjdrew/gosproto

# Features
Similar to [golang protobuf](https://github.com/golang/protobuf) version 3

No need use sproto.Int32( ), sproto.String() wrapper to generate point field in order to implement optional field

Optional will be auto processed by comparing incoming value and type default value

	bool default is false

	integer default is 0
	
	string default is ""

Fields equals to its default value will be skipped encoding and decoding automatically'

USE SPROTO LIKE NORMAL STRUCT!

# Type map
sproto type      | golang type
---------------- | -------------------------------------------------
string           | \string, []byte
integer          | \int8, \uint8, \int16, \uint16, \int32, \uint32, \int64, \uint64, \int, \uint
boolean          | \bool
object           | \struct
*string    		 | []string
*integer 		 | []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []int, []uint
*boolean         | []bool
*struct          | []\*struct
*struct(index)   | []\*struct

# Schema parser and meta info

https://github.com/davyxu/gosproto/tree/master/meta

# Code generator

https://github.com/davyxu/gosproto/tree/master/sprotogen

## Install

```
	go get -u -v github.com/davyxu/gosproto/sprotogen
```

## Usage

```
	# generate go source file
	sprotogen --type=go --out=addressbook.go --gopackage=example test.sproto
	
```

# Exsample

https://github.com/davyxu/gosproto/tree/master/example

```golang
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
	}

	var sample AddressBook
	_, err = sproto.Decode(data, &sample)

	if err != nil {
		t.Log(err)
	}

```


# Test

```golang
go test github.com/davyxu/gosproto
```


# Benchmark


```
$ go test -bench . github.com/davyxu/gosproto
BenchmarkEncode-8         	  500000	      2498 ns/op
BenchmarkDecode-8         	  500000	      3134 ns/op
BenchmarkEncodePacked-8   	  500000	      2894 ns/op
BenchmarkDecodePacked-8   	  500000	      3480 ns/op
PASS
ok  	github.com/davyxu/gosproto	6.162s
```

* xjdrew/gosproto Version

```
$ go test -bench . github.com/xjdrew/gosproto
BenchmarkEncode-8         	 1000000	      1931 ns/op
BenchmarkDecode-8         	  500000	      3498 ns/op
BenchmarkEncodePacked-8   	  500000	      2476 ns/op
BenchmarkDecodePacked-8   	  500000	      3896 ns/op
PASS
ok  	github.com/xjdrew/gosproto	7.024s
```




