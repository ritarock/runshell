package action

import (
	"runshell/lib/command"
	"runshell/lib/file"
	rs "runshell/lib/string"
	"strings"
)

func Run(path string) {
	for _, commandStrList := range file.ReadLines(path) {
		for _, commandStr := range rs.CreateCmd(commandStrList) {
			for _, v := range commandStr {
				command.Run(strings.Split(v, ","))
			}
		}
	}
}
