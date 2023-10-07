package basics

import "fmt"

// ArrayUsage
/*
声明数组
var arrayName [size]dataType
var numbers = [5]int
声明并初始化数组
var numbers = [5]int{1, 2, 3, 4, 5}
*/
func ArrayUsage() {
	intArrayTest()
	multidimensionalArrayTest()
}

func intArrayTest() {
	var nums [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		nums[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("nums[%d] = %d\n", j, nums[j])
	}

	for j := range nums {
		fmt.Printf("nums[%d] = %d\n", j, nums[j])
	}
}

func multidimensionalArrayTest() {
	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中.不会补齐长度
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("animals[%d]: %v, lenth of row:%d\n", i, animals[i], len(animals[i]))
		for j := range animals[i] {
			fmt.Printf("animals[%d][%d]: %v\n ", i, j, animals[i][j])
		}
	}
}
