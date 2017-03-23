// extend standard


message PhoneNumber {

	// 头注释
	number string // 尾注释

    // 整形
	type int32
		
}


message Person {
	
	name string
		
	id int32
		
	email string
		
	phone []PhoneNumber
		
}


message AddressBook {
	
	person []Person
		
}


enum MyCar {
	
	Monkey = 1
		
	Monk = 2
		
	Pig = 3
		
}


message MyData {
	
	name  string	
		
	type  MyCar	
		
	int32  int32    // extend standard

	uint32  uint32	
		
	int64  int64	
		
	uint64  uint64	
		
	bool  bool	
		
}


message MyProfile {
	
	nameField MyData
		
	nameArray []MyData
		
	nameMap []MyData(type)
		
}


