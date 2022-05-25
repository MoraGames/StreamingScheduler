package password

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type SHA3_512Password []byte

func NewSHA3_512Password(password []byte) SHA3_512Password {
	//Generate the hash
	hashedPassword := sha3.Sum512(password)

	//Generate the password
	return SHA3_512Password(hashedPassword[:])
}

func (hash SHA3_512Password) ToString() string {
	return hex.EncodeToString(hash)
}