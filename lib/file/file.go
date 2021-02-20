package file

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {
	var result []string
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		result = append(result, s.Text())
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	return result
}
