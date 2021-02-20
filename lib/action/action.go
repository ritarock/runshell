package action

import (
	"fmt"
	"runshell/lib/command"
	"runshell/lib/file"
	rs "runshell/lib/string"
	"strings"
)

func Run(path string) {
	readCommandsList := strings.Split(file.Read(path), "\n")

	for _, commmandList := range readCommandsList {
		for _, commandStr := range rs.CreateCmd(commmandList) {
			for _, v := range commandStr {
				fmt.Println(v)
				go command.Run(strings.Split(v, ","))
			}
		}
	}
}
