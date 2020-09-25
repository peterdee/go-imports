package main

import (
	"fmt"
)

// this function is private, and it's accessible only inside the 'main' package
func logger(value string) {
	fmt.Println(value)
}
