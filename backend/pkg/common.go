package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"time"
)

const (
	SessionExpireAfter = time.Hour * 24 * 7 // default is 7 days
)

func GenerateUUID() string {
	return uuid.New().String()
}
func Hash(mainStr string, saltStr string) string {
	bytes := sha256.Sum256([]byte(mainStr + "_" + saltStr))
	return hex.EncodeToString(bytes[:])
}

func HashUserPassword(password string, userId string) string {
	return Hash(password, userId)
}
