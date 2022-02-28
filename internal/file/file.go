package file

import (
	"bufio"
	"os"
)

func Read(path string) ([]string, error) {
	result := []string{}
	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			result = append(result, fileScanner.Text())
		}
	}
	if err := fileScanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
