package main

import "fmt"

func main() {
	const (
		a = 10
		b = "a"
		c = iota
	)

	var (
		d = 10
	)
	fmt.Println("a =", a, "b =", b, "c=", c)
	fmt.Println("d=", d)
}