package utils

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"slogv2/src/main/utils/customError"
	"strconv"
	"time"
)

func ScryptPassword(password string) (string, string, error) {
	const KeyLen = 32

	//salt := TestDefaultSalt
	//使用时间戳作为salt
	var salt []byte
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < KeyLen; i++ {
		salt = append(salt, byte(rand.Intn(10)))
	}

	N := 16384
	R := 8
	P := 1

	hash, err := scrypt.Key([]byte(password), salt, N, R, P, KeyLen)
	if err != nil {
		return "", "", customError.GetError(customError.OTHER_ERROR, "scrypt password failed"+err.Error())
	}

	pwd := base64.StdEncoding.EncodeToString(hash)
	strSalt := ""
	for _, v := range salt {
		strSalt += strconv.Itoa(int(v))
	}

	return pwd, strSalt, nil
}

func ScryptPasswordBySalt(password string, salt string) (string, error) {
	const KeyLen = 32

	N := 16384
	R := 8
	P := 1

	//!! salt需要手动转换为[]byte
	//注意此处逻辑
	//加密时使用的salt为生成的32位int串一位位压入
	//解密时重建salt应将string类型的salt一位位转换为int类型压入
	var saltByte []byte
	for _, v := range salt {
		num, _ := strconv.Atoi(string(v))
		saltByte = append(saltByte, byte(num))
	}

	hash, err := scrypt.Key([]byte(password), saltByte, N, R, P, KeyLen)
	if err != nil {
		return "", customError.GetError(customError.OTHER_ERROR, "scrypt password failed"+err.Error())
	}

	pwd := base64.StdEncoding.EncodeToString(hash)

	return pwd, nil
}

func CheckPassword(password string, storedPwd string, salt string) (int, error) {
	pwd, err := ScryptPasswordBySalt(password, salt)

	fmt.Printf("get pwd: %s\n", pwd)
	fmt.Printf("get storedPwd: %s\n", storedPwd)
	fmt.Printf("salt : %s\n", salt)

	if err != nil {
		return customError.FAIL, err
	}

	if pwd != storedPwd {
		return customError.FAIL, nil
	}

	return customError.SUCCESS, nil
}
