package main

import (
	"fmt"
	"time"
)

/**
 * 结构体：结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体
 */
func main()  {
	// 结构体定义
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}
	type x struct {
		a,b,c int
	}
	d := x{
		a: 1,
		b: 2,
		c: 3,
	}
	// 结构体赋值
	a := Employee{
		ID:        1,
		Name:      "段晋",
		Address:   "成都市高新区",
		DoB:       time.Now(),
		Position:  "Jug",
		Salary:    0,
		ManagerID: 0,
	}
	a.Salary = 100
	fmt.Printf("结构体a的数据为：%v\n", a)
	// 对成员取地址，通过指针访问
	b := &a.Address
	c := "我的地址是" + *b
	fmt.Println(c)
	fmt.Printf("结构体d的数据为：%v\n", d)
}
