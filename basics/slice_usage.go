package basics

import "fmt"

//SliceUsage
/**
切片是对数组的抽象。数组的长度不可改变，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大.
切片声明（声明未指定长度的数组）：
var slice1 []type
或者通过内置的make函数声明切片（通过len指定切片初始大小）
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)
切片初始化： slice1 :=[] int {1,2,3,4,5}
也可以在make函数中指定容量，即切片最大长度
slice1 := make([]T, length, capacity)
*/
func SliceUsage() {

	/*声明切片： var slice [] int*/
	var slice []int
	fmt.Printf("slice: %v, capacity of slice: %d, length of slice: %d\n", slice, cap(slice), len(slice))

	/*切片声明及初始化： var slice1 = []int{1, 2, 3, 4, 5} 或 slice1 :=[] int {1,2,3,4,5}*/
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1: %v, capacity of slice1: %d, length of slice1: %d\n", slice1, cap(slice1), len(slice1))

	/*make函数声明切片： slice2 := make([]T, length, capacity), 初始长度的所有元素为对应类型的默认空值, 只能通过给对应下表赋值去修改 slice[index]=2*/
	slice2 := make([]int, 5, 10)

	/*切片 slice2 = slice追加slice1全部元素*/
	slice2 = append(slice, slice1...)
	fmt.Printf("slice2: %v, capacity of slice2: %d, length of slice2: %d\n", slice2, cap(slice2), len(slice2))
	/*切片 slice2 追加切片 slice1 某个区间(slice1[startIndex:endIndex])的所有元素(...表示所有元素)(不限制数量，返回新的切片)*/
	slice3 := append(slice, slice1[0:3]...)
	fmt.Printf("slice3: %v, capacity of slice3: %d, length of slice3: %d\n", slice3, cap(slice3), len(slice3))

	/* 创建切片 slice4 是切片slice1的两倍容量*/
	slice4 := make([]int, len(slice1)*2, (cap(slice1))*2)
	/*将slice1 复制到 slice4 ,从slice1下标0到结束一一复制，slice4其余初始长度下表元素为默认空值*/
	copy(slice4, slice1)
	fmt.Printf("slice4: %v, capacity of slice4: %d, length of slice4: %d\n", slice4, cap(slice4), len(slice4))

	slice5 := make([]int, len(slice1)*2, (cap(slice1))*2)
	/*将slice 复制到 slice5 ,从slice下标0到结束一一复制，slice5其余初始长度下表元素为默认空值*/
	copy(slice5, slice)
	fmt.Printf("slice5: %v, capacity of slice5: %d, length of slice5: %d\n", slice5, cap(slice5), len(slice5))
}
