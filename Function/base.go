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
	// 对函数值的调用
	e := map[int]string{0: "aaa", 1: "ccc"}
	f := addMap
	fmt.Printf("调用函数值的结果为：%v\n", f(2, "abc", e))
	g := squares()
	fmt.Println(g()) // "1"dd
	fmt.Println(g()) // "4"
	fmt.Println(g()) // "9"
	fmt.Println(g()) // "16"
	fmt.Printf("可变参数的函数结果为：%v\n", sum(2, 3, 4, 5, 6))
}

/**
 * 函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
 */
func addMap(a int, b string, c map[int]string) map[int]string {
	c[a] = b
	return c
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// 可变参数
func sum(a int, vals ...int) int  {
	total := a
	for _, v := range vals {
		total += v
	}
	// defer延迟函数 defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后
	defer test(total)
	return total
}

func test(a int)  {
	fmt.Println(22222222222)
}