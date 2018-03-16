package main

import (
	"fmt"
	"worker"
)

func main() {

	num := 0
    Loop:fmt.Println("请输入LOFTER或tumblr用户名或豆瓣相册编号：\n" +
		"例如：lofter username或者douban 1234567890")
	var site, name string
	if num > 3 {
		fmt.Println("输入已达上限")
	}

	fmt.Scan(&site, &name);

	fmt.Println("输入", site, name)

	switch site {
	case "lofter":
		go worker.Lofter(name)
	case "douban":
		go worker.Douban(name)
	case "tumblr":
		go worker.Tumblr(name)
	default:
		fmt.Println("Please enter right address!")
	}
	num += 1
	goto Loop
}
