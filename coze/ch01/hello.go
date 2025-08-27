package ch01

import "fmt"

func Hello() {
	fmt.Println("hello gfw")
}

func Run_slice() {
	var s = []int{4, 1, 2, 3}
	fmt.Println(s[4:])
}
