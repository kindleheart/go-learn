package main

import (
	"errors"
	"fmt"
)

type myError struct {
}

func (e myError) Error() string {
	return "this is my error"
}

func main() {
	bottomErr := &myError{}

	// %w 包装错误
	middleErr := fmt.Errorf("this is a middle error, %w", bottomErr)

	// Unwrap 接触包装，返回被包装的 error
	if errors.Unwrap(middleErr) == bottomErr {
		fmt.Println("unwrap")
	}

	topErr := fmt.Errorf("this is a top error, %w", middleErr)
	fmt.Printf("%+v\n", topErr)

	// Is 判断是不是错误链上的err
	if errors.Is(topErr, bottomErr) {
		fmt.Println("is")
	}

	// As 类型转换为错误链上特定的error
	be := &myError{}
	if errors.As(topErr, &be) {
		fmt.Printf("%+v\n", be)
	}

	// Json 包装多个错误
	fmt.Println("---------")
	err := errors.Join(bottomErr, middleErr, topErr)
	fmt.Println(err)

	Delay()
}

func Delay() {
	fns := make([]func(), 0, 10)

	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
