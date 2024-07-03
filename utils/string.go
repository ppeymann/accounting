package utils

import (
	"crypto/aes"
	"crypto/cipher"
	cRand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrSecret = errors.New("secret is not in correct length")
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// RandNumberDigits this function for create random string of table values with specific length
func RandNumberDigits(length int) string {
	bytes := make([]byte, length)
	n, err := io.ReadAtLeast(cRand.Reader, bytes, length)
	if n != length {
		log.Println("this is error for randDigit", err)
	}

	for i := 0; i < len(bytes); i++ {
		bytes[i] = table[int(bytes[i])%len(table)]
	}

	return string(bytes)
}

// HashString this function is for hashing input string
// and returns hashed string of using golang bcrypt algorithm.
func HashString(val string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(val), 10)
	return string(bytes), err
}

// CheckHashedString compare a hashed string with specific plain string.
func CheckHashedString(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}

// EncryptText encrypts given string by specific secret.
func EncryptText(value, secret string) (string, error) {
	text := []byte(value)
	key := []byte(secret)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrSecret
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cRand.Reader, nonce); err != nil {
		return "", err
	}

	res := gcm.Seal(nonce, nonce, text, nil)
	b64 := base64.URLEncoding.EncodeToString(res)

	return b64, nil
}

// DecryptText decrypt given cipher string by specific secret.
func DecryptText(value, secret string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	key := []byte(secret)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", nil
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
