package common

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v2/log"
)

func GenerateSecretKey(length int) string {
	randombytes := make([]byte, length)
	_, err := rand.Read(randombytes)
	if err != nil {
		log.Infof("Error while generate randombytes: ", err)
	}

	key := base64.StdEncoding.EncodeToString(randombytes)

	return key
}
