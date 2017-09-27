# gosproto
[sproto](https://github.com/cloudwu/sproto)云风的sproto协议描述文件及代码生成工具

English Doc see below:
https://github.com/davyxu/gosproto/blob/master/README_en.md

# 代码基于

https://github.com/xjdrew/gosproto

# 特性

* 兼容云风版sproto格式,并扩充的格式,比Protobuf更方便, 更快

* 代码生成输出支持go,c#,lua以及云风原版sproto

# 描述文件格式(*.sp)

下面是sprotogen支持的推荐格式
```

enum Vocation {
	Monkey
	Monk
	Pig
}

message PhoneNumber {

	number string

	type int32
}


message Person {

	name string

	id  int32

	email string

	phone PhoneNumber

	voc Vocation
}

message AddressBook {

	person []Person
}


```
## 特性

* 自动生成tag序列号(base0,sproto规范),也可以手动指定

* 自动生成枚举序号(base0),也可以手动指定

* 类go结构体字段命名方式

* 比Protobuf更方便的导出注释内容做协议扩充

## 兼容云风的sproto描述格式(*.sproto)

解析器可以同时兼容两种格式书写,可以在任何地方混合书写. 更推荐使用sp格式,有更强的扩展和兼容性

```
.PhoneNumber {

	number 0 : string

	type 1 : integer

}

.Person {

	name 0 : string

	id 1 : integer

	email 2 : string

	phone 3 : *PhoneNumber

}

.AddressBook {

	person 0 : *Person

}
```


# Go版的使用方法
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

# 支持类型

* int32: 32位整形
* int64: 64位整形
* uint32: 无符号32位整形
* uint64: 无符号64位整形
* string: 字符串
* float32: 单精度浮点数(默认精度0.001)
* float64: 双精度浮点数(默认精度0.001)
* bytes: 二进制数据
* enum: int32封装
* bool: 布尔
* message 结构体

所有类型前添加[]表示数组

# 编译

```
	go get -u -v github.com/davyxu/gosproto/sprotogen
```

# 下载

https://github.com/davyxu/gosproto/releases

# sprotogen命令行参数

* go_out
    输出Go源码

* lua_out
    输出lua源码,兼容云风版本

* cs_out
    输出C#源码, 配套基类库请参考https://github.com/davyxu/sproto-Csharp

* sproto_out
    输出云风版sproto的描述文件

* package
    包或者命名空间名

* cellnet_reg
    在Go的生成代码中添加cellnet自动注册入口

# 使用方法

```
	# 输出go源码
	sprotogen --go_out=addressbook.go --package=example addressbook.sp
	
	# 生成C#源码
	sprotogen --cs_out=addressbook.cs --package=example addressbook.sp
	
```

# 例子

https://github.com/davyxu/gosproto/tree/master/example



# 可选工具:Protobuf描述格式转sproto描述格式

https://github.com/davyxu/gosproto/tree/master/pb2sproto

## 特性
保留所有protobuf的注释

## 安装方法

```
	go get -u -v github.com/davyxu/gosproto/pb2sproto
```
第三方库依赖: github.com/davyxu/pbmeta

## 使用方法

```
	# 使用Protobuf编译器:protoc配合插件github.com/davyxu/pbmeta/protoc-gen-meta 生成pb的上下文信息
	# 参见: github.com/davyxu/gosproto/pb2sproto/Make.bat
	
	# 根据上下文信息导出sproto
	pb2sproto --pbmeta=meta.pb --outdir=.	
	
```

# 备注

感觉不错请star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/sunicdavy

提交bug及特性: https://github.com/davyxu/gosproto/issues
