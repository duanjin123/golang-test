package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 获取标准输入
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("循环获取标准输入：%s\n", arg)
	}
	// 直接打印标准输入，获取到切片slice
	fmt.Println(args)
	// 字符长度
	fmt.Println(len(args[1]))
	// 标准输入连接字符串
	for _, arg := range args {
		arg += " |"
		fmt.Printf("连接后的字符串为：%s\n", arg)
	}
	fmt.Printf("连接后的字符串为：%s", strings.Join(args, " |"))
	// 空map类型定义
	m := make(map[string]int)
	m["key1"] = 1
	fmt.Printf("\n%v\n", m)
	// 打印请求输出
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Printf("request %s: return: %v", "https://duanjin.life", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("request %s: return: %v", "https://duanjin.life", err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf("request %s: return: %s", "https://duanjin.life", b)
	// web服务器
	http.HandleFunc("/", getPath)
	http.ListenAndServe("localhost:4000", nil)
}

func getPath(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(r, "请求的路径为：%s", req.URL.Path)
}
