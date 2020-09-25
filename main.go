package main

import (
	"go-imports/level1"                 // the name of the imported package is 'level1'
	Cool "go-imports/level1/level2"     // 'Cool' is an alias for 'level2'
	. "go-imports/level1/level2/level3" // Imports with the dot are special
)

func main() {
	logger("Level Zero")
	level1.LevelOne("Level One: easy")
	Cool.LevelTwo("Level Two: well yea")

	// 'level3' was imported with the dot, so its contents are available in the 'main' scope
	// but 'level3' private stuff IS NOT available
	LevelThree("Level Three: lets go deeper")
}
