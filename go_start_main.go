/*
package为 main 时，才能识别main() 方法作为程序启动入口
使用 go run go_start_main.go 运行代码
使用 go build go_start_main.go 编译代码成为可执行文件
*/
package main

/**
import 引包
import "fmt"
import "time"
简写形如下：
*/
import (
	"fmt"
	"go-demo/basics"
	"time"
)

/*
*
func 定义函数
main函数为程序启动入口
函数名字首字母小写为private私函数
函数名字首字母大写为public公有函数
编码默认一行为一句, 当一行写入多条语句，需要人为使用;隔开
*/
func main() {
	// 标准输出
	fmt.Println("Hello, Welcome to Go world!")
	// 格式化输出
	fmt.Printf("Today is:%s\n", time.Now().Format(time.DateOnly))
	fmt.Printf("NowTime is:%s\n", time.Now().Format(time.TimeOnly))
	fmt.Printf("NowTime is:%s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	// 输入输出测试
	//basics.InputOutput()
	// 变量测试
	fmt.Printf("var Age in basics: %d\n", basics.Age)
	fmt.Printf("var Lang in basics: %s\n", basics.Lang)
	basics.VarTest()
	basics.VarTest2()
	basics.VarTest3()
	// 常量测试
	fmt.Printf("const AGE in basics: %d\n", basics.AGE)
	fmt.Printf("const LANG in basics: %s\n", basics.LANG)
	basics.ConstTest()
	basics.ConstTest2()
	basics.ConstTest3()
	// 类型测试
	fmt.Printf("AbsoluteZeroC:%2.f\n", basics.AbsoluteZeroC)
	var tempC = basics.AbsoluteZeroC
	var f = float64(tempC)
	fmt.Println(f == float64(tempC))
	fmt.Println(0 <= float64(tempC))
	fmt.Printf("tempC to f:%.2f\n", f)
	var tempF = basics.CToF(tempC)
	fmt.Printf("AbsoluteZeroC to tempF:%.2f\n", tempF)
	var tempC0 = basics.FToC(0)
	fmt.Printf("tempF0 to tempC0:%.2f\n", tempC0)

	//运算符测试
	basics.MathOperator()
	basics.LogicOperator()
	basics.RelationalOperator()

	//条件语句测试
	//basics.SwitchTest()
	//basics.IfElseTest()

	//函数测试
	basics.FuncUsage()

	// 派生类型， 除bool、number(int,float)、string等基本类型外，都是派生类型，包含以下类型：
	// 指针类型(Pointer)
	basics.PointerUsage()
	// 数组类型(array)
	basics.ArrayUsage()
	// 切片类型 (slice)
	basics.SliceUsage()
	// 结构化类型(struct)
	basics.StructUsage()
	// Map类型(map)
	// Channel类型 (chan)
	// 函数类型(func)
	// 接口类型(interface)

}
