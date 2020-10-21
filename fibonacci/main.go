package main

import "fmt"

func fibonacci(f int64) int64 {
	if f > 1 {
		return (fibonacci(f-1) + fibonacci(f-2))
	} else {
		if f == 1 {
			return 1
		}
		if f == 0 {
			return 0
		}
		return -1
	}
}

func main() {
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(12))
	fmt.Println(fibonacci(15))
}
