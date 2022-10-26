package main

import "fmt"

// 批量声明
var (
	name  string = "小白"
	age   uint
	isBoy bool
)

// 常量
const time = 1
const (
	n1 = iota //0　iota出现，是常量的行索引，注意是行索引，一行多个也是一行，从０开始
	n2 = 100  //100
	n3 = iota //2
	n4        //3 不写跟上一行一致
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	str := "类型推导"
	_ = "匿名变量"
	fmt.Println(str)
	fmt.Println("Hello World")
}
