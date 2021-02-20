package string

import (
	"regexp"
)

func CreateCmd(str string) [][]string {
	re := regexp.MustCompile(`[^[].*[^]]`)

	res := re.FindAllStringSubmatch(str, -1)

	return res
}
