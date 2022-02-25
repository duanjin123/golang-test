package main

import "fmt"

func main()  {
	a := 2
	b := "test"
	c := map[int]string {
		0: "aaa",
		1: "bbb",
	}
	d := addMap(a, b, c)
	fmt.Printf("map c经过函数addMap后的值为：%v\n", d)
}

/**
 * 函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
 */
func addMap(a int, b string, c map[int]string) map[int]string {
	c[a] = b
	return c
}
