package main

import "fmt"

/**
 * slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的底层引用了一个数组对象
 * slice包含指针、长度、容量三部分
 */
func main()  {
	// 切片赋值（对数组的引用）
	i := [...]int{1,2,3,4,5,6,7,8,9}
	i1 := i[1:4]
	i2 := i[3:6]
	i3 := i[:4]
	fmt.Print(i1)
	fmt.Print(i2)
	fmt.Print(i3)
	fmt.Printf("i1的指针为%v\n", &i1)
	fmt.Printf("i1的长度为%d\n", len(i1))
	fmt.Printf("i1的容量为%d\n", cap(i1))
	fmt.Printf("i1的第4个元素为：%s\n", i1[3])
	// slice翻转
	k1 := [...]int{11,21,31,41,51,61,71}
	fmt.Print(reverse2(&k1))
	fmt.Print(rotate(k1[:], 2))
	//fmt.Print(reverse1(j[:7]))
	//// 向slice追加元素
	//k := j[:5]
	//k = append(k, 3)
	//fmt.Print(k)
	k2 := [...]string{"aa", "bb", "bb", "a1", "b2", "aa", "aa", "a1"}
	fmt.Print(eliminat(k2[:]))
}

// slice翻转
func reverse1(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 1. 重写reverse函数，使用数组指针代替slice
func reverse2(s *[7]int) *[7]int  {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}

// 2.编写一个rotate函数，通过一次循环完成旋转。
func rotate(s []int, n int) []int {
	for i := 0;i < n; i++ {
		s = append(s, s[0])
		s = s[1:len(s)]
	}
	return s
}

// 3. 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func eliminat(string1 []string) []string {
	for i,j := 0,1; j < len(string1); i,j = i+1, j+1 {
		if string1[i] == string1[j] {
			copy(string1[j:], string1[j+1:])
			string1 = string1[:len(string1) - 1]
			i,j = i-1,j-1
		}
	}
	return string1
}