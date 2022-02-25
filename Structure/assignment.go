package main

import "fmt"

type structA struct {
	A int
	B string
	C structB
}

type structB struct {
	AA int
	BB string
}

func main()  {
	// int类型赋值
	var a int
	fmt.Println(IntAssign(&a))
	// 字符串赋值
	var b string = "aaa"
	c := &b
	b = "bbb"
	fmt.Println(b)
	// c变量指向变量b的地址
	// 打印出c的值
	fmt.Println(*c)
	// 结构体赋值
	var d structA
	var e structB
	d.A = 1
	d.B = "B"
	e.AA = 2
	e.BB = "BB"
	d.C = e
	fmt.Print(d)
	// map
	m := make(map[string]int)
	v1 := m["a"]                // map查找，失败时返回零值
	v2 := &m
	fmt.Print(v1)
	fmt.Print(*v2)
	//v3 := <-ch
}

func IntAssign(a *int) int {
	*a = 2
	return *a
}
