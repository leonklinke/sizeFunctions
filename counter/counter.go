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

	methodCounts := fileAnalyzer(file)

	fmt.Println("Methods: ", len(methodCounts))
	fmt.Println("Avg size: ", avg(methodCounts))

	return count
}

func fileAnalyzer(file *os.File) []int {
	scanner := bufio.NewScanner(file)

	methodsCount := methodCounter(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return methodsCount
}

func methodCounter(scanner *bufio.Scanner) []int {
	methodLineCounter := make([]int, 0)
	methodID := 0
	openningBraces := 0
	closingBraces := 0
	openedBrace := false
	firstLine := false

	for scanner.Scan() {
		openning, closing := countBraces(scanner.Text())

		if openning > 0 && !openedBrace {
			methodLineCounter = append(methodLineCounter, 0)
			openedBrace = true
			firstLine = true
		}

		if !openedBrace {
			continue
		}

		if openning > closing {
			openningBraces += openning - closing
		}
		if openning < closing {
			closingBraces += closing - openning
		}

		if openningBraces == closingBraces {
			methodID++
			openningBraces = 0
			closingBraces = 0
			openedBrace = false
			continue
		}

		if firstLine {
			firstLine = false
			continue
		}

		methodLineCounter[methodID]++
	}

	return methodLineCounter
}

func countBraces(line string) (int, int) {
	openningBraces := 0
	closingBraces := 0
	for _, char := range line {
		switch char {
		case '{':
			openningBraces++
		case '}':
			closingBraces++
		}
	}
	return openningBraces, closingBraces
}

func avg(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}
	return sum / len(slice)
}
