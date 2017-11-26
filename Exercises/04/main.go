package main

import (
	"fmt"
)

func main() {
	var smallnum int
	var bignum int
	fmt.Println("Hi there, please enter a small number:")
	fmt.Scan(&smallnum)
	fmt.Println("and a bigger number:")
	fmt.Scan(&bignum)
	i := bignum % smallnum
	fmt.Println("Remainder:", i)
}
