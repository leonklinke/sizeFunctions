package main

import (
	"github.com/leonklinke/sizeFunctions/counter"
)

var lines int

func main() {

	printWelcome()
	// linesThreshold := getLinesThreshold()
	// linesThreshold := 20
	// getProjectPath()
	projectPath := "./probeFunctions"
	lines = counter.CountProjectFunctionLines(projectPath)
}
