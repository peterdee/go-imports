package level2

import (
	"fmt"
)

// this function is private, and it's accessible only inside the 'level2' package
func privateLogger(value string) {
	fmt.Println(value)
}

// this function is public, it's accessible if package is imported in other places
func LevelTwo(value string) {
	privateLogger(value)
}
