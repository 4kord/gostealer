package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

func DecryptWithKey(key []byte, text []byte) string {

	//remove v10
	ciphertext := text[3:]
	//
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}

	return string(plaintext)
}
