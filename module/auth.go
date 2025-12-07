package module

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"music-twin-backend/config"
	"strconv"
	"strings"
	"time"
)

func GetUserIdFromToken(ctx context.Context, token string) int32 {
	if token == "" {
		return 0
	}

	output, err := Decrypt(token)
	if err != nil {
		return 0
	}
	splitStrings := strings.Split(output, "|")
	if len(splitStrings) != 2 {
		return 0
	}

	userIdRaw, err := strconv.ParseInt(splitStrings[0], 10, 0)
	if err != nil {
		return 0
	}

	return int32(userIdRaw)
}

func GenerateUserToken(ctx context.Context, userId int) string {
	currentTime := time.Now().Format("2006-01-02 15:04:05") //.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")

	stringText := fmt.Sprint(userId, "|"+currentTime)

	token, err := encrypt(stringText)
	if err != nil {
		return ""
	}

	return token
}

func encrypt(input string) (cipherString string, err error) {

	text := []byte(input)
	var block cipher.Block
	var ciphertext []byte

	key := []byte(config.GetHashSecretKey())
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	ciphertext = make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	cipherString = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

func Decrypt(token string) (plaintext string, err error) {

	ciphertext, _ := base64.StdEncoding.DecodeString(token)
	key := []byte(config.GetHashSecretKey())
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
