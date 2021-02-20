package file

import (
	"io/ioutil"
	"log"
	"os"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
