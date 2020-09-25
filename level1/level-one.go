package level1

import (
	"fmt"
)

// this function is private, and it's accessible only inside the 'level1' package
func privateLogger(value string) {
	fmt.Println(value)
}

// this function is public, it's accessible if package is imported in other places
func LevelOne(value string) {
	privateLogger(value)
}
