package main

import "fmt"

func intercambia(a *uint, b *uint) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	var a, b uint
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	intercambia(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
