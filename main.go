package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	age  int
	name string
}

func main() {
	u := User{11, "xixi发撒放假啊萨弗拉斯发了啥打法放大发发啊是的"}
	fmt.Println(unsafe.Sizeof(u))
	fmt.Println(unsafe.Sizeof(u.age), unsafe.Sizeof(u.name))
	str := "agafasdf"
	fmt.Println(unsafe.Sizeof(str))
}
