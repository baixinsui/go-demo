package basics

type Celsius float64    // 定义类型摄氏温度,底层数据类型使用数据类型
type Fahrenheit float64 // 定义类型华氏温度,底层数据类型使用数据类型

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度,初始化使用数据类型赋值
	FreezingC             = 0       // 结冰点温度
	BoilingC              = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // 摄氏温度可以使用其底层数据类型的运算和操作,也可以与底层数据类型直接转换
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) // 华氏温度可以使用其底层数据类型的运算和操作,也可以与底层数据类型直接转换
}
