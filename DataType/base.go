package main

import (
	"fmt"
	"strconv"
)

/**
 * 基础数据类型
 */
func main()  {
	intFuc()
	floatFuc()
	complexFuc()
	boolFuc()
	stringFuc()
	constFuc()
	trans()
}

/**
 * 整型
 */
func intFuc()  {
	// 定义
	var a int32
	// 赋值
	a = 64
	fmt.Printf("int类型a的值为：%d\n", a)
	b := &a
	*b += 24
	fmt.Println("int类型b+24的值为：%d\n", *b)
	*b *= 3
	fmt.Println("int类型b*3的值为：%d\n", *b)
	// 位运算
	c := int32(1)
	d := c >> 2 | 4
	fmt.Println("1左移2位与4做或运算的值为：%d\n", d)
}

func floatFuc()  {
	// 定义
	var a float32 = 23.45
	fmt.Println("float类型a的值为：%d\n", a)
	a += 32.663
	fmt.Println("float类型a+32.663的值为：%d\n", a)
}

func complexFuc()  {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println("复数类型x的值为：%d\n", x)
	fmt.Println("复数类型y的值为：%d\n", y)
	z := x*y
	fmt.Println("x*y值为：%d\n", z)
}

func boolFuc()  {
	var a bool
	fmt.Println("bool类型a的值为：%d\n", a)
	// 逻辑运算中 &&比||优先级更高
}

func stringFuc()  {
	// 定义
	var a string = "hello, world!"
	fmt.Println("string类型a的值为：%d\n", a)
	// 字符长度
	length := len(a)
	fmt.Println("a的长度为：%d\n", length)
	// 索引查找时返回字符的字节码
	fmt.Println("string类型a[2]的值为：%d\n", a[2])
	// 截取
	b := a[0:5]
	fmt.Println("a前5个字符的值为：%d\n", b)
	// 字符串连接符 +
	c := a[:5] + " lincoo"
	fmt.Println("c的值为：%d\n", c)
	// 通过字节码生成字符串
	d := string(92)
	fmt.Println("字节码为92的字符为：%s\n", d)
}

func constFuc()  {
	const name = "aaa"
	// 批量声明常量时，除了第一个，其余可以省略初始化表达式，省略表达式后，表示使用前面相同的表达式进行初始化
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Printf("name常量的值为：%s\n", name)
	fmt.Printf("d常量的值为：%s\n", d)
	// iota常量生成器
	const (
		Sunday = iota // 第一个初始化为0，后续每个加1，类似于枚举
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("Sunday常量的值为：%s\n", Sunday)
	fmt.Printf("Friday常量的值为：%s\n", Friday)
}

/**
 * 类型转换
 */
func trans()  {
	// 数字转换为字符串
	var a int32 = 43
	x := fmt.Sprintf("%s", a)
	x += "aaa"
	fmt.Printf("a转换为字符串为：%s\n", x)
	// bool转换为字符串
	y := strconv.FormatBool(1 == 1)
	fmt.Printf("a转换为bool为：%v\n", y)
}