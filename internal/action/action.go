package action

import (
	"fmt"
	"log"
	"ritarock/runshell/internal/command"
	"ritarock/runshell/internal/file"
	"sync"
)

func asyncRun(commands string, isFin chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	result, err := command.Run(commands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	isFin <- true
}

func Run(path string) error {
	commandsList, err := file.Read(path)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	isFin := make(chan bool, len(commandsList))

	for _, commands := range commandsList {
		wg.Add(1)
		go asyncRun(commands, isFin, &wg)
	}
	wg.Wait()
	close(isFin)

	return nil
}
