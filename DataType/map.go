package main

import "fmt"

/**
 * 一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。
 */
func main()  {
	// map的定义
	var a map[string]int
	b := map[string]string {
		"apple": "two",
		"orange": "three",
		"egg": "twenty",
		"test": "test",
	}
	c := make(map[string]int)
	// map赋值
	b["apple"] = "five"
	c["a"] = 1
	fmt.Printf("map a的值为：%v\n", a)
	fmt.Printf("map b的值为：%v\n", b)
	fmt.Printf("map c的值为：%v\n", c)
	// 删除map元素
	delete(b, "test")
	fmt.Printf("map b删除元素后的值为：%v\n", b)
	// map循环
	for key, value := range b {
		fmt.Printf("map b中" + key + "对应的值为：%s\n", value)
	}
	
}