package main

import (
	"fmt"
)

func printWelcome() {
	fmt.Println("------------------------------")
	fmt.Println("-- Welcome to Size function --")
	fmt.Println("--    Let's clean it up     --")
	fmt.Println("------------------------------")
	fmt.Println("")
}

func getLinesThreshold() int {
	gotLines := false

	for !gotLines {
		lines, err := ascLinesThreshold()

		if err != nil {
			fmt.Println(err)
			continue
		}
		return lines
	}

	return 0
}

func ascLinesThreshold() (int, error) {
	var lines int
	fmt.Println("")
	fmt.Println("Please, enter the number of lines you want as a limit (30 recomended) ")
	fmt.Print("-> ")

	fmt.Scanf("%d", &lines)
	if lines <= 0 {
		return 0, fmt.Errorf("Wrong value")
	}

	return lines, nil
}
