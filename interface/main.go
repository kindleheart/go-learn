package main

import "fmt"

type A interface {
	a()
	b()
}

type B interface {
	a()
	c()
}

type C interface {
	A
	B
}

type D struct {
}

func (d D) a() {
	fmt.Println("a")
}

func (d D) b() {
	fmt.Println("b")
}
func (d D) c() {
	fmt.Println("c")
}

func main() {
	var a C = D{}
	a.a()
	fmt.Println("aaa")
}
