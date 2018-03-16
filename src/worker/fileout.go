package worker

import (
	"os"
	"fmt"
)

func GetFile(filename string) *os.File{
	var f *os.File
	var err error
	filename = filename + ".txt"
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
	}
	check(err)
	return f
}




func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}