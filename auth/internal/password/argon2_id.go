package password

import (
	"crypto/rand"
	"log"

	"golang.org/x/crypto/argon2"
)

type ARGON2_IDPassword struct {
	salt []byte
	hash []byte
}

func NewARGON2_IDPassword(password []byte) ARGON2_IDPassword {
	//Generate the salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Panicln(err)
	}

	//Generate the hash
	hash := argon2.IDKey(password, salt, 1, 64*1024, 4, 64)

	//Generate the password
	return ARGON2_IDPassword{salt, hash}
}

func (password ARGON2_IDPassword) ToString() string {
	return string(password.salt) + string(password.hash)
}