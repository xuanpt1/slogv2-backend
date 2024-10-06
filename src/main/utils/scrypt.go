package utils

import (
	"golang.org/x/crypto/scrypt"
	"slogv2/src/main/utils/customError"
)

func ScryptPassword(password string) (string, string, error) {
	const KeyLen = 32

	salt := TestDefaultSalt

	N := 16384
	R := 8
	P := 1

	hash, err := scrypt.Key([]byte(password), []byte(salt), N, R, P, KeyLen)
	if err != nil {
		return "", "", customError.GetError(customError.OTHER_ERROR, "scrypt password failed"+err.Error())
	} else {
		return string(hash), salt, nil
	}
}
