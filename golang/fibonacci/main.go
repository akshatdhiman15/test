package main

import "fmt"

func fib(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

func main() {

	a := 0
	b := 1
	n := 10
	for i := 0; i < n; i++ {
		c := a + b
		fmt.Println(b)
		a = b
		b = c
	}

	v := fib(10)

	fmt.Println("10th fib:", v)
}
