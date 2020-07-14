package main

import "fmt"

/**
包声明
引入包
函数
变量
语句 & 表达式
注释
*/

//同一个目录下面不能有个多 package main


func main() {
	fmt.Println("Hello, World!")

	// 声明一个变量并初始化
	var a = "Hello Tom"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	/**
	数值类型（包括complex64/128）为 0
	布尔类型为 false
	字符串为 ""（空字符串）
	以下几种类型为 nil：

	var a *int
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error // error 是接口
	*/

	var i int
    var f float64
    var bo bool
    var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)
	


	//省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	ff := "Runoob" // var f string = "Runoob"
    fmt.Println(ff)


	/**
	*多变量声明
	*/

	//类型相同多个变量, 非全局变量
	var x01, y01 int
	// 这种因式分解关键字的写法一般用于声明全局变量
	var (  
		a01 int
		b01 bool
	)

	var c01, d01 int = 1, 2
	var e01, f01 = 123, "hello"

	//这种不带声明格式的只能在函数体中出现
	//g, h := 123, "hello"

	g01, h01 := 123, "hello"
    println(x01, y01, a01, b01, c01, d01, e01, f01, g01, h01)

}
