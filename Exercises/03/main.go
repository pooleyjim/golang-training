package main

import (
	"fmt"
)

func main() {
	// Print to terminal asking for name.
	var name string
	fmt.Println("Please enter name:")
	fmt.Scan(&name)
	fmt.Println("Hello ", name, " with ", name, ".")
}
