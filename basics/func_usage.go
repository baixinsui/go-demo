package basics

import "fmt"

// FuncUsage
/* 函数定义
func function_name( [parameter list] ) [return_types] {
   函数体
}
*/
func FuncUsage() {
	fmt.Printf("max(2,6)=%v\n", max(2, 6))
	fmt.Println(swap("John", "Hi"))

	/* 定义局部变量 */
	var a = 100
	var b = 200
	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 使用值传递调用函数来交换值 */
	fmt.Println("使用值传递调用函数来交换值")
	swapInt(a, b)
	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	/* 使用引用传递调用函数来交换值
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	fmt.Println("使用引用传递调用函数来交换值")
	swapRefer(&a, &b)

	fmt.Printf("交换后 &a 的值 : %d\n", a)
	fmt.Printf("交换后 &b 的值 : %d\n", b)
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}

}

// 返回多值
func swap(x, y string) (string, string) {
	return y, x
}

// 传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
// 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
func swapInt(x, y int) int {
	var temp int
	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/
	return temp
}

/* 定义交换值函数,引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数*/
func swapRefer(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
