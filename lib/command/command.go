package command

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func Run(commandList []string) {
	AsyncRun(commandList)
}

func AsyncRun(commandList []string) {
	var wg sync.WaitGroup

	for _, command := range commandList {
		wg.Add(1)
		go func(command string) {
			out, err := exec.Command("sh", "-c", command).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(out))
			wg.Done()
		}(command)
		wg.Wait()
	}
}
