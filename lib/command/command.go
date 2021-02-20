package command

import (
	"fmt"
	"log"
	"os/exec"
)

func Run(commandList []string) {
	for _, command := range commandList {
		out, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(string(out))
	}
}
