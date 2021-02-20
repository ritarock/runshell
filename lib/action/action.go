package action

import (
	"runshell/lib/command"
	"runshell/lib/file"
	rs "runshell/lib/string"
	"strings"
	"sync"
)

func Async(commandList string, isFin chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, commandStr := range rs.CreateCmd(commandList) {
		for _, v := range commandStr {
			command.Run(strings.Split(v, ","))
		}
	}
	isFin <- true
}

func Run(path string) {
	wg := new(sync.WaitGroup)
	readCommandsList := strings.Split(file.Read(path), "\n")
	isFin := make(chan bool, len(readCommandsList))

	for _, commandlist := range readCommandsList {
		wg.Add(1)
		go Async(commandlist, isFin, wg)
	}
	wg.Wait()
	close(isFin)
}
