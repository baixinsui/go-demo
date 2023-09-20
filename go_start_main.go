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
	"baixinsui/go-demo/basics"
	"fmt"
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

	// 输入输出测试
	basics.InputOutput()
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
}