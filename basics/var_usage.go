// Package basics
package basics

// Lang 声明变量并初始化赋值&类型推断
var Lang = "English"

// Age 声明变量并初始化赋值
var Age int = 20

func VarTest() {
	// 声明变量
	var name string
	var age int
	var isMarried bool
	var weight float32
	// 输出各自类型的初始值
	PrintInfo(name, age, isMarried, weight, Lang)
	// 变量赋值
	name = "Tom"
	age = 20
	isMarried = false
	weight = 55.50
	PrintInfo(name, age, isMarried, weight, Lang)
}

// VarTest2
func VarTest2() {
	// 短变量赋值,只能用在函数内
	name := "Tom2"
	age := 30
	isMarried := true
	// 类型推断默认float64
	weight := 50.50
	PrintInfo(name, age, isMarried, float32(weight), Lang)
}

func VarTest3() {
	// 批量初始化变量
	var name, age, isMarried, weight = "Tom3", 32, false, 58.35
	PrintInfo(name, age, isMarried, float32(weight), Lang)
}
