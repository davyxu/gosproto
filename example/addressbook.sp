

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

.MyCar {
	Monkey 	1
	Monk 	2
	Pig 	3
}


.MyData{
	name 0: string
	type 1: MyCar
	int32 2: int32	# extend standard
}

.MyProfile{
	nameField 1: MyData
	nameArray 2: *MyData
	nameMap 3: *MyData(type)
}