package level3

import (
	"fmt"
)

// this function is private, it's accessible only in the scope of 'level3'
// and it's NOT AVAILABLE if imported with the dot
func privateLevel3Logger() {
	fmt.Println("Shoot me down, but I won't fall, I am titanium")
}

// this function is public, it's accessible if package is imported in other places
func LevelThree(value string) {
	fmt.Println(value)
	privateLevel3Logger()
}
