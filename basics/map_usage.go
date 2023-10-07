package basics

import "fmt"

// MapUsage
/*
Map 是一种无序的键值对的集合,通过 key 来快速检索数据，key 类似于索引，指向数据的值。
在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。
Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

可以使用内建函数 make 或使用 map 关键字来定义 Map:
map_variable := make(map[KeyType]ValueType, initialCapacity)
其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。
Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，Map 会自动扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。
或
m := map[string]int
*/
func MapUsage() {

	/*make(map[KeyType]ValueType, initialCapacity) 声明map变量*/
	var countryCapitalMap = make(map[string]string, 7)
	fmt.Printf("Before initial countryCapitalMap: %v, lenth:%d\n", countryCapitalMap, len(countryCapitalMap))
	/*初始化map变量*/
	countryCapitalMap = map[string]string{
		"China":       "Beijing",
		"USA":         "Washington",
		"UK":          "London",
		"Japan":       "Tokyo",
		"South Korea": "汉城",
		"North Korea": "Pyongyang",
	}
	fmt.Printf("After initial countryCapitalMap: %v, lenth:%d\n", countryCapitalMap, len(countryCapitalMap))
	/*map 添加键值对或修改key值*/
	countryCapitalMap["Russia"] = "Moscow"
	countryCapitalMap["Mexico"] = "Mexico City"
	countryCapitalMap["South Korea"] = "Seoul"
	fmt.Printf("After modify countryCapitalMap: %v, lenth:%d\n", countryCapitalMap, len(countryCapitalMap))

	fmt.Printf("The capital of country %s is %s\n", "China", countryCapitalMap["China"])
	fmt.Printf("The capital of country %s is %s\n", "French", countryCapitalMap["French"])

	/*遍历Map, range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value
	*/
	for k, v := range countryCapitalMap {
		fmt.Printf("The capital of country %s is %s\n", k, v)
	}
	/*只遍历key, 获取map的key set*/
	var keys []string
	for key := range countryCapitalMap {
		keys = append(keys, key)
		fmt.Printf("The capital of country %s is %s\n", key, countryCapitalMap[key])
	}
	fmt.Printf("The keys of map is %v\n", keys)

	/*只遍历value, 获取map的value set*/
	var values []string
	for _, value := range countryCapitalMap {
		values = append(values, value)
	}
	fmt.Printf("The values of map is %v\n", values)

	/*短赋值定义并初始化map*/
	userAgeMap := map[string]int{
		"John": 16,
		"Yarn": 15,
		"Slim": 17,
	}
	fmt.Printf("userAgeMap: %v\n", userAgeMap)
	/*获取元素*/
	fmt.Printf("The age of user %s is %d\n", "French", userAgeMap["French"])
	/* 如果键不存在，ok 的值为 false，对应的值为该类型的零值*/
	age, ok := userAgeMap["John"]
	fmt.Printf("Get the age of user %s result: %v, age is %d\n", "John", ok, age)
	age2, ok2 := userAgeMap["Yarn2"]
	fmt.Printf("Get the age of user %s result: %v, age is %d\n", "Yarn2", ok2, age2)

	/*删除键值对*/
	delete(userAgeMap, "Slim")
	fmt.Printf("After delete userAgeMap: %v\n", userAgeMap)

}
