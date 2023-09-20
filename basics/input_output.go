// Package basics
package basics

import (
	"fmt"
)

func InputOutput() {
	// 测试标准输入及标准输出
	// 定义变量接收输入值,Go中命名区分大小写
	var name string
	var age int
	var weight float32
	var isMarried bool
	fmt.Println("Enter your name:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		println(err)
		fmt.Printf("Input value of name:%v is invalid!", &name)
	}
	fmt.Println("Enter your age:")
	_, err = fmt.Scanln(&age)
	if err != nil {
		println(err)
		fmt.Printf("Input value of age:%v is invalid!", &age)
	}
	fmt.Println("Enter your weight:")
	_, err = fmt.Scanln(&weight)
	if err != nil {
		println(err)
		fmt.Printf("Input value of weight:%v is invalid!", &weight)
	}
	fmt.Println("Are you married:")
	_, err = fmt.Scanln(&isMarried)
	if err != nil {
		println(err)
		fmt.Printf("Input value of isMarried:%v is invalid!", &isMarried)
	}
	// 格式化输出
	fmt.Printf("name:%s,", name)
	fmt.Printf("age:%d,", age)
	fmt.Printf("isMarried:%v,", isMarried)
	fmt.Printf("weight:%v,", weight)
	fmt.Printf("language:%v\n", Lang)
	// Sprintf格式化并返回值
	var info = SprintInfo(name, age, isMarried, weight, "")
	fmt.Sprintln(info)
}

func PrintInfo(name string, age int, isMarried bool, weight float32, lang string) {
	fmt.Printf("name:%v,", name)
	fmt.Printf("age:%v,", age)
	fmt.Printf("isMarried:%v,", isMarried)
	fmt.Printf("weight:%v,", weight)
	if lang == "" {
		fmt.Printf("language:%s\n", LANG)
	} else {
		fmt.Printf("language:%s\n", lang)
	}
}

func SprintInfo(name string, age int, isMarried bool, weight float32, lang string) string {
	if lang == "" {
		lang = Lang
	}
	return fmt.Sprintf("name:%v,age:%v,isMarried:%v,weight:%v,language:%s", name, age, isMarried, weight, lang)
}
