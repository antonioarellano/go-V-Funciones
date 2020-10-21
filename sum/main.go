package main

import "fmt"

func sum(args ...uint) uint {
	m := args[0]
	for _, v := range args {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	fmt.Println(sum(155, 4, 7, 3, 55, 44, 7, 75))
	fmt.Println(sum(4, 44, 7, 75))
	fmt.Println(sum(4, 7, 3, 55, 44, 7, 75, 545, 3, 66))
}
