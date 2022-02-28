package main

/**
 * 在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。
 */

import "fmt"

type P struct { p,q int }

type Q struct {
	x P
	y int
	z int
}

type A int

// 函数
func sum(p, q P) int  {
	return p.p + q.q + p.q + q.p
}

// 方法
func (p P) sum(q P) int  {
	return p.p + q.q + p.q + q.p
}

func (a A)sum2(b A) int  {
	return int(a + b)
}

func (p *P) sum3(q *P) int  {
	return  p.p + q.q + p.q + q.p
}

func main()  {
	var p P =  P{ 1, 2 }
	var q P =  P{ 3, 4 }
	a := sum(p, q)
	fmt.Printf("函数sum计算的结果是：%d\n", a)
	// 这儿的p叫做方法的接收器，p.sum的表达式叫做选择器
	b := p.sum(q)
	fmt.Printf("方法sum计算的结果是：%d\n", b)
	var c A = 2
	d := c.sum2(3)
	fmt.Printf("方法sum2计算2，3的结果是：%d\n", d)
	e := p.sum3(&q)
	fmt.Printf("方法sum3计算的结果是：%d\n", e)
	fmt.Printf("方法sum3计算的结果是：%d\n", (&P{1, 2}).sum3(&q))
	q1 := Q{
		x: P{1, 2},
		y: 2,
		z: 3,
	}
	fmt.Printf("通过嵌入结构体的值是：%v\n", q1)
}
