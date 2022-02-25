package main

import (
	"fmt"
	"strconv"
)

/**
 * 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
 */
func main()  {
	// 数组定义
	var a [3]int
	b := [4]string{"a", "b", "c", "d"}
	c := [...]int{1,2,3}
	fmt.Printf("数组a为：%v\n", a)
	fmt.Printf("数组b为：%v\n", b)
	fmt.Printf("数组c为：%v\n", c)
	// 定义51个元素的数组，最后一个元素为1
	h := [...]int{50:1}
	h[23] = 2
	fmt.Print(h)
	// 数组元素获取
	fmt.Printf("\n数组b第2个元素为：%s\n", b[1])
	b[1] = "bb"
	fmt.Printf("修改后数组b第2个元素为：%s\n", b[1])
	// 数组循环
	for i, v := range c {
		fmt.Printf("数组c第" + strconv.Itoa(i+1) + "个元素为：%s\n", v)
	}
}
