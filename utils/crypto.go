package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"passmanager/config"

	"github.com/mergermarket/go-pkcs7"
	logger "github.com/sirupsen/logrus"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// encrypt encrypts plain string with a secret key and returns encrypt string.
func Encrypt(plainData string, passphrase string) (string, error) {
	cipherBlock, err := aes.NewCipher([]byte(createHash(passphrase)))
	if err != nil {
		logger.Error("Error in NewCipher : ", err)
		return "", err
	}
	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		logger.Error("Error in NewGCM : ", err)
		return "", err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Error("Error in ReadFull : ", err)
		return "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil)), nil
}

// decrypt decrypts encrypt string with a secret key and returns plain string.
func Decrypt(encodedData string, passphrase string) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		logger.Error("Decrypt: Error in DecodeString. Error: ", err)
		return "", config.ErrInvalidHashKey
	}
	cipherBlock, err := aes.NewCipher([]byte(createHash(passphrase))) //aes.NewCipher(secret)
	if err != nil {
		return "", err
	}
	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}
	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}
	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func SimpleEncrypt(text string) (string, error) {

	fmt.Println("text--->>", text)
	block, err := aes.NewCipher([]byte(os.Getenv("EMAIL_ENC_KEY")))
	if err != nil {
		return "", err
	}
	plaintext := []byte(text)
	fmt.Println(plaintext)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Println("encrypted--->>>")
	return encodeBase64(ciphertext), nil
}

func SimpleDecrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("EMAIL_ENC_KEY")))
	if err != nil {
		return "", err
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encryption(unencrypted string) (string, error) {
	CIPHER_KEY := os.Getenv("ENC_KEY")
	key := []byte(CIPHER_KEY)
	plainText := []byte(unencrypted)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decryption(encrypted string) (string, error) {
	CIPHER_KEY := os.Getenv("ENC_KEY")
	key := []byte(CIPHER_KEY)
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%s", cipherText), nil
}
