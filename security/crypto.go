// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"io"

	"github.com/pkg/errors"
)

const (
	MATTERMOST_VER						 = "9745736730483574"
)

func EncryptionKey() (result string) {
		result = MATTERMOST_VER
		return
	}

func Encrypt(str string) (encryptedStr string, err error) {
	CIPHER_KEY := []byte(MATTERMOST_VER)
	
	if(!strings.HasPrefix(str, "ENC")) {
		encryptedStr, err = encryptString(CIPHER_KEY, str)
		encryptedStr = "ENC" + encryptedStr
	} else {
		encryptedStr = str
	}
	return

}

func encryptString(key []byte, message string) (encmess string, err error) {
		plainText := []byte(message)

		block, err := aes.NewCipher(key)
		if err != nil {
			return
		}

		//IV needs to be unique, but doesn't have to be secure.
		//It's common to put it at the beginning of the ciphertext.
		cipherText := make([]byte, aes.BlockSize+len(plainText))
		iv := cipherText[:aes.BlockSize]
		if _, err = io.ReadFull(rand.Reader, iv); err != nil {
			return
		}

		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

		//returns to base64 encoded string
		encmess = base64.URLEncoding.EncodeToString(cipherText)
		return
	}


func decryptString(key []byte, securemess string) (decodedmess string, err error) {
		cipherText, err := base64.URLEncoding.DecodeString(securemess)
		if(err!=nil) {
			return
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			return
		}

		if len(cipherText) < aes.BlockSize {
			err = errors.New("Ciphertext block size is too short!")
			return
		}

		//IV needs to be unique, but doesn't have to be secure.
		//It's common to put it at the beginning of the ciphertext.
		iv := cipherText[:aes.BlockSize]
		cipherText = cipherText[aes.BlockSize:]

		stream := cipher.NewCFBDecrypter(block, iv)
		// XORKeyStream can work in-place if the two arguments are the same.
		stream.XORKeyStream(cipherText, cipherText)

		decodedmess = string(cipherText)
		return
	}


func Decrypt(str string) (decryptedStr string, err error) {
	CIPHER_KEY := []byte(MATTERMOST_VER)
	
	if(strings.HasPrefix(str, "ENC")) {
		decryptedStr = strings.TrimPrefix(str, "ENC")
		decryptedStr, err = decryptString(CIPHER_KEY, decryptedStr)
	} else {
		decryptedStr = str
	}

	return
}