package crypt

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	memory      uint32 = 64 * 1024
	time        uint32 = 2
	parallelism uint8  = 2
	saltLength  uint32 = 16
	keyLength   uint32 = 32
)

var (
	ErrInvalidEncodedHash          = errors.New("the encoded hash in not valid")
	ErrIncompatibleArgon2IDVersion = errors.New("incompatible version of argon2id")
	ErrPasswordsDoNotMatch         = errors.New("password do not match")
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := generateFromPassword(password)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, encodedHash string) (err error) {
	salt, hash, err := extractSaltAndHash(encodedHash)
	if err != nil {
		return
	}

	comparableHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLength)

	if subtle.ConstantTimeCompare(hash, comparableHash) != 1 {
		return ErrPasswordsDoNotMatch
	}

	return
}

func generateFromPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// $algorithm$version$params<memory,time,parallelism>$salt$hash
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateSalt() ([]byte, error) {
	bytes := make([]byte, saltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func extractSaltAndHash(encodedHash string) (salt []byte, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, ErrInvalidEncodedHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrIncompatibleArgon2IDVersion
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, err
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, err
	}

	return
}
