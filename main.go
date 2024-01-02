package main

import (
	"fmt"
	"os"
)

func main() {
	stat, err := os.Stat("file.go")
	if err != nil {
		return
	}
	fmt.Println(stat.Name(), stat.Mode())
}
