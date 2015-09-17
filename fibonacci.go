package main

import "fmt"

func main() {
	// fmt.Println(fib(55))
	
	for a:=0; a<10; a++ {
		fmt.Println(fib(a))
	}
}

func fib(i int) int {
	if i==0 {
		return 0
	}
	if i==1 {
		return 1
	}
	
	return fib(i-1) + fib(i-2)
}