package main

import (
	"./counter"
	"./vender"
	"flag"
	"fmt"
	"time"
)

var path string

func init() {
	flag.StringVar(&path, "d", "", "请输入计算的文件")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("用法：counter -d xx.txt")
	}
}

func main() {
	fmt.Println("计算开始(Ctrl C 退出)")
	_, err := vender.Exists(path)
	if err != nil {
		panic("不存在该文件或目录")
	}
	if vender.IsDir(path) {
		//todo:
	} else {
		counter.DoSingle(path)
	}
	for ; ; {
		time.Sleep(1)
	}
}
