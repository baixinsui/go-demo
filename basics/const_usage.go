// Package basics
package basics

// LANG 声明常量时必须初始化,常量一般使用全大写
const LANG = "Chinese"

// AGE 声明常量时必须初始化,常量一般使用全大写
const AGE int = 20

func ConstTest() {
	// 声明常量时必须初始化,常量一般使用全大写
	const NAME string = "Tom"
	const AGE int = 30
	const IsMarried bool = true
	const WEIGHT float32 = 55.50
	// 常量值不能修改
	//AGE=36
	// 输出各自类型的初始值
	PrintInfo(NAME, AGE, IsMarried, WEIGHT, "")
}
func ConstTest2() {
	// 常量类型推断
	const NAME = "Tom2"
	const AGE = 32
	const IsMarried = true
	// 类型推断默认float64
	const WEIGHT = 50.50
	// 常量值不能修改
	//AGE=36
	PrintInfo(NAME, AGE, IsMarried, float32(WEIGHT), "")
}

func ConstTest3() {
	// 常量批量赋值&类型推断
	const NAME, AGE, IsMarried, WEIGHT = "Tom3", 33, false, 55.50
	// 常量值不能修改
	//AGE=36
	PrintInfo(NAME, AGE, IsMarried, float32(WEIGHT), LANG)
}
