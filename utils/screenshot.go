package utils

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/kbinani/screenshot"
)

func GetScreenshot(logFolderPath string) {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Println(err)
			return
		}
		fileName := fmt.Sprintf("DISPLAY_%d.png", i)
		file, _ := os.Create(path.Join(logFolderPath, fileName))
		defer file.Close()
		png.Encode(file, img)
		log.Printf("[SCREEN] Took screenshot of %d screen\n", i)
	}
}
