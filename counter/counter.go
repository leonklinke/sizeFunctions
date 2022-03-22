package counter

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countFunctionLines(fullPath string) int {
	var count int
	file, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
