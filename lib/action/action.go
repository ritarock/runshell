package action

import (
	"runshell/lib/command"
	"runshell/lib/file"
	rs "runshell/lib/string"
	"strings"
	"sync"
)

func asyncRun(commandList string, isFin chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, commandStr := range rs.CreateCmd(commandList) {
		for _, v := range commandStr {
			command.Run(strings.Split(v, ","))
		}
	}
	isFin <- true
}

func Run(path string) {
	var wg sync.WaitGroup
	readCommandsList := strings.Split(file.Read(path), "\n")
	isFin := make(chan bool, len(readCommandsList))

	for _, commandlist := range readCommandsList {
		wg.Add(1)
		go asyncRun(commandlist, isFin, &wg)
	}
	wg.Wait()
	close(isFin)
}
