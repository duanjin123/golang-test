package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 获取命令行输入
	fmt.Println(getArgs())
	fmt.Println(loop())
	fmt.Println(join())
	// 直接打印，获取输入参数的切片
	fmt.Println(os.Args[1:])
}

func getArgs() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func loop() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func join() string {
	return strings.Join(os.Args[1:], " ")
}
