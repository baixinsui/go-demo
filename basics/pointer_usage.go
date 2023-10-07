package basics

import "fmt"

// PointerUsage
/*
指针声明格式如下：
var var_name *var-type
一个指针变量指向了一个值的内存地址; 给指针变量赋值，实际变量类型必须与指针变量类型一致，在实际变量前加 & 前缀
在变量前面加上 & 号来获取相应变量的内存地址；
在指针类型前面加上 * 号来获取指针所指向的内容；
修改指针指向的变量的实际值后，指针指向值也会变化
*/
func PointerUsage() {
	var a = 20  /* 声明实际变量 */
	var ip *int /* 声明指针变量 */
	ip = &a     /* 指针变量的存储地址 */
	fmt.Printf("a 变量的值是: %d\n", a)
	fmt.Printf("a 变量的地址是: %d\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %d\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	/*修改变量值*/
	a = 30
	fmt.Printf("修改a值后 a 变量的值是: %d\n", a)
	fmt.Printf("修改a值后 a 变量的地址是: %d\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("修改a值后 ip 变量储存的指针地址: %d\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("修改a值后 *ip 变量的值: %d\n", *ip)

}
