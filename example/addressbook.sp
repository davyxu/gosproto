

.PhoneNumber {
	number 0 : string
	type 1 : int32
}
	
	
.Person {
    name 0 : string
    id 1 : int32
    email 2 : string



    phone 3 : *PhoneNumber
}

.AddressBook {
    person 0 : *Person
}


# Full sytax of sproto

enum MyCar {
	Monkey 	= 1
	Monk 	= 2
	Pig 	= 3
}


.MyData{
	name 0: string
	type 1: MyCar
	int32 2: int32	# extend standard
	uint32 4: uint32
	int64 5: int64
	uint64 6: uint64
}

.MyProfile{
	nameField 1: MyData
	nameArray 2: *MyData
	nameMap 3: *MyData(type)
}