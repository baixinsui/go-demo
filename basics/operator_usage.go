package basics

import "fmt"

// MathOperator 算术运算符
func MathOperator() {

	var a = 21
	var b = 10
	var c int
	//双目运算符
	//加法
	c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	//减法
	c = a - b
	fmt.Printf("%d - %d = %d\n", a, b, c)
	//乘法
	c = a * b
	fmt.Printf("%d * %d = %d\n", a, b, c)
	//整除(结果符号与被除数相同,除数不能为0)
	c = a / b
	fmt.Printf("%d / %d = %d\n", a, b, c)
	fmt.Printf("%d / %d = %d\n", -5, -2, -5/-2)
	fmt.Printf("%d / %d = %d\n", -5, 2, -5/-2)
	//求余 (%%转义%)
	c = a % b
	fmt.Printf("%d %% %d = %d\n", a, b, c)
	// 单目运算法
	//自增, a=a+1
	fmt.Printf("a=%d, ", a)
	a++
	fmt.Printf("a++, a=%d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	//自减, a=a-1
	fmt.Printf("a=%d, ", a)
	a--
	fmt.Printf("a--, a=%d\n", a)
}

// LogicOperator 逻辑运算符
func LogicOperator() {

	var a = true
	var b = false
	var c bool
	c = a && b
	fmt.Printf("%v && %v = %v\n", a, b, c)
	c = a || b
	fmt.Printf("%v || %v = %v\n", a, b, c)
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	c = a && b
	fmt.Printf("%v && %v = %v\n", a, b, c)
	c = a || b
	fmt.Printf("%v || %v = %v\n", a, b, c)
	c = !(a && b)
	fmt.Printf("!(%v && %v) = %v\n", a, b, c)
}

// RelationalOperator 关系运算符
func RelationalOperator() {
	var a = 21
	var b = 10
	var c bool
	c = a == b
	fmt.Printf("%d == %d : %v\n", a, b, c)

	c = a != b
	fmt.Printf("%d != %d : %v\n", a, b, c)

	c = a > b
	fmt.Printf("%d > %d : %v\n", a, b, c)

	c = a < b
	fmt.Printf("%d < %d : %v\n", a, b, c)

	c = a >= b
	fmt.Printf("%d >= %d : %v\n", a, b, c)

	c = a <= b
	fmt.Printf("%d <= %d : %v\n", a, b, c)

}
