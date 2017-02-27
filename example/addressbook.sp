// extend standard


message PhoneNumber {
	# 头 
	number 0 string# 哈哈		
		# 头 
# 哈哈	
# 头
# 哈哈 
	type 1 int32	
		
}


message Person {
	
	name 0 string	
		
	id 1 int32	
		
	email 2 string	
		
	phone 3 *PhoneNumber	
		
}


message AddressBook {
	
	person 0 *Person	
		
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
	
	nameField 1 MyData	
		
	nameArray 2 *MyData	
		
	nameMap 3 *MyData(type)	
		
}


