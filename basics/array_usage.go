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
	var nums [10]int /* nums 是一个长度为 10 的数组 */
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
	/*num2 声明并初始化，初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	注意：数组中，数组长度是必要的要素，长度不同的两个数组即使元素类型相同，两个变量的类型也不同， 即变量 num2 与 num 是不同的类型
	*/
	var nums2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2)

	/*短赋值变量nums3，nums3 与nums2 数组元素类型和长度都相同，两者类型才相同 */
	nums3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums3)

	/*[...]可用于推断数组长度 */
	var nums4 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0, 556.00}
	fmt.Println(nums4)

	/*使用数组指定区间的所有元素初始化另一个数组: [0:2] 左闭右开即起始下标为0,结束下标为2-1， 结束下标不能大于数组长度*/
	nums5 := nums2[0:2]
	fmt.Println(nums5)
	/*使用数组指定区间的所有元素初始化另一个数组: [2:5] 左闭右开即起始下标为2,结束下标5-1*/
	nums6 := nums2[2:5]
	fmt.Println(nums6)

	/*使用数组指定区间的所有元素初始化另一个数组:
	[startIndex:] 左闭右开即起始下标为startIndex,省缺的结束下标数组最后元素的下标
	[:endIndex] 左闭右开,省缺的开始下标数组第一个元素的下标即0，结束下标为startIndex
	[:] 左闭右开,开始、结束都省缺，即取全部元素，相当于  nums9 := nums2
	*/
	nums7 := nums2[0:]
	fmt.Println(nums7)
	nums8 := nums2[:3]
	fmt.Println(nums8)
	nums9 := nums2[:]
	fmt.Println(nums9)
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
