package main

import (
	"fmt"
)

func set_zero() {
	//1 &^ 0
	//1 &^ 1 ---0
	//0 &^ 1 ---0
	//0 &^ 0 ---0
	//fmt.Println(^-2)   //1 : -2, 2: -3,  -2: 1   ^符号直接使用为取反-1
	//1     :0000 0000 0000 0001
	//-2buma:1111 1111 1111 1110  符号位不变，取反+1
	//res   :0000 0000 0000 0001  按位与结果为0
	//-2    :1000 0000 0000 0010
	fmt.Println(1 & -2)
	fmt.Println(1 &^ 0) //---1
	fmt.Println(1 &^ 1) //---0
	fmt.Println(0 &^ 1) //---0
	fmt.Println(0 &^ 0) //---0
	var tmp int = 10
	tmp_ := &tmp
	fmt.Println(tmp_)  // 变量地址
	fmt.Println(*tmp_) // 指针变量 由内存地址指向变量值
}
func set_permission() {
	// 定义用户权限
	const (
		Readable   int = 1 << iota // 1
		Writeable                  // 2
		Executable                 //4
	)
	// 校验用户权限 perm
	var perm int
	perm = 7                                   // 具有Readable+Writeable+Executable权限
	fmt.Println(perm&Readable == Readable)     // true
	fmt.Println(perm&Writeable == Writeable)   // true
	fmt.Println(perm&Executable == Executable) // true
	perm = 6                                   // 具有Writeable+Executable权限
	fmt.Println(perm&Readable == Readable)     // false
	fmt.Println(perm&Writeable == Writeable)   // true
	fmt.Println(perm&Executable == Executable) // true
	perm = 3                                   // 具有Readable+Writeable权限
	fmt.Println(perm&Readable == Readable)     // true
	fmt.Println(perm&Writeable == Writeable)   // true
	fmt.Println(perm&Executable == Executable) // false
	perm = 0                                   // 无权限
	fmt.Println(perm&Readable == Readable)     // false
	fmt.Println(perm&Writeable == Writeable)   // false
	fmt.Println(perm&Executable == Executable) // false
	// 清除用户对应权限 用按位 置零
	perm = 7
	fmt.Println(perm)
	fmt.Println(perm&^Readable == Readable)     // false  清除读权限
	fmt.Println(perm&^Writeable == Writeable)   // false  清除写权限
	fmt.Println(perm&^Executable == Executable) // false  清除执行权限
}

func main() {
	set_zero()
	//set_permission()
}
