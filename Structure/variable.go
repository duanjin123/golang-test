package main

import "fmt"

var a int

func main()  {
	var x,y int = 1, 2
	p := &x
	fmt.Printf("变量x的指针地址为：%s\t\n", &x)
	fmt.Printf("变量y的指针地址为：%s\t\n", &y)
	fmt.Printf("变量x的指针的值为：%s\t\n", *p)
	fmt.Printf("变量x的值为：%s\t\n", x)
	v := 1
	fmt.Println(&v)
	fmt.Println(f() == f())
	fmt.Println(incr(&a))
	fmt.Println(incr(&a))
	fmt.Println(incr(&a))
	// new函数创建变量，返回指针地址
	b := new(int)
	*b = 2
	fmt.Println(*b)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int  {
	*p++
	return *p
}

// 总结
// 1. 变量的定义一般用var x int = 2（即：var 变量 类型 表达式组成，类型或表达式可以省略），也可简化为x := 2
// 2. x := 2中，&x表示该变量保存的指针地址，*x表示指针指向的地址的值，也可以通过*x = 3修改变量的值
// 3. 每次定义变量，都会生成一个新的指针地址
// 4. 垃圾回收机制：变量的生命周期取决于指针或引用是否可达，变量存储与堆/栈取决于是否逃逸，全局变量必须在堆上分配，局部变量可在堆/栈上进行分配，逃逸的变量需要额外的内存分配，因此可能会降低程序性能