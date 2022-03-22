package counter

import (
	"io/ioutil"
	"log"
)

func CountProjectFunctionLines(projectPath string) int {
	var count int
	files, err := ioutil.ReadDir(projectPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			count = countFunctionLines(projectPath + "/" + file.Name())
		}
	}

	return count
}
