package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"slogv2/src/main/utils/customError"
	"strings"
	"time"
)

//DONE 生成JWT，保存登录状态

// Claims 自定义声明结构体并内嵌jwt.StandardClaims
type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func ReleaseToken(userId uint) (string, error) {
	//创建一个我们自己的声明
	expiresAt := time.Now().Add(JWTDefaultExpire).Unix()
	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "xuanpt2",
			Subject:   "user token",
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString([]byte(JWTDefaultSecret))
	if err != nil {
		return "", customError.GetError(customError.JWT_CREATE_ERROR, err.Error())
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	//解析token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTDefaultSecret), nil
	})
	if err != nil {
		return nil, nil, customError.GetError(customError.JWT_VERIFY_ERROR, err.Error())
	}

	return token, claims, nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenUserId(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)

	token, claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, customError.GetError(customError.JWT_VERIFY_ERROR, err.Error())
	}
	if !token.Valid {
		return 0, customError.GetError(customError.JWT_VERIFY_ERROR, "token valid failed")
	}
	return claims.UserId, nil
}

func VerifyToken(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTDefaultSecret), nil
	})
	if err != nil {
		return customError.GetError(customError.JWT_VERIFY_ERROR, err.Error())
	}
	return nil
}
