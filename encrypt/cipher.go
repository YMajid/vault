package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

/*
Encrypt accepts a key and plaintext and returns an hex representation version of the encrypted value.

Code is based on the standard library example at:
	- https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
*/
func Encrypt(key string, plaintext string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

/*
Decrypt accepts a key and a cipherHex (hex representation of ciphertext) and decrypt it.

Code is based on the standard library example at:
	- https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
*/
func Decrypt(key string, cipherHex string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hash := md5.New()
	fmt.Fprint(hash, key)
	cipherKey := hash.Sum(nil)

	return aes.NewCipher(cipherKey)
}
