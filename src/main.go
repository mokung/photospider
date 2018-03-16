package main

import (
	"fmt"
	"worker"
)

var cnum chan int

func main() {

	num := 0
	cnum = make(chan int, 3)
Loop:
	fmt.Println("请输入LOFTER或tumblr用户名或豆瓣相册编号：\n" +
		"例如：lofter username或者douban 1234567890")
	var site, name string

	fmt.Scan(&site, &name);
	fmt.Println("输入", site, name)

	switch site {
	case "lofter":
		go worker.Lofter(name, cnum)
	case "douban":
		go worker.Douban(name, cnum)
	case "tumblr":
		go worker.Tumblr(name, cnum)
	default:
		fmt.Println("Please enter right address!")
		num -= 1
	}
	num += 1
	if num > 3 {
		fmt.Println("输入已达上限")
		for i := 0; i < num; i++ {
			<-cnum
		}
		fmt.Println("WE DONE!!!")
		num = 0
	}
	goto Loop
}
