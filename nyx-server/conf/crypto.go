package conf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func IsGoodClient(ip string) bool {
	panel, err := GetPanel()
	if err != nil {
		log.Printf("Error getting panel config: %v", err)
		return false
	}

	rvg := panel.Admins[ip].RegisteredViaGood

	return rvg
}

func PassSalt() string {
	hasher := sha256.New()
	hasher.Write([]byte(MachinePassword + Saltstr))
	Saltstr = Saltgen()

	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}

func VerifyPassword(Hash string) bool {
	expectedHash := PassSalt()

	Saltstr = Saltgen()

	return Hash == expectedHash
}

func Saltgen() string {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(salt)
}
