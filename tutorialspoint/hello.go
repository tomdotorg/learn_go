package main

import (
	"fmt"
	"log"
)

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	log.Printf("%s", "hello, world!")
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {

	var foo, bar int32
	var s = "string"
	foo = 5
	fmt.Println("foo", foo)
	fmt.Println(bar)
	fmt.Println(&s)
	//	fmt.Println(*s)
	/* This is my first sample program. */
	_foo := 5
	_foo += 5
	fmt.Println("Hello, World!")
	fmt.Println(max(2, 3))
	fmt.Println(swap("second", "first"))
	fmt.Println("I am in Go Programming World!")
	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)
	}
}
