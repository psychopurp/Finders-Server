package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret_key = "elyar"

type TokenKeys map[string]interface{}

//生成Token，并把keys 放进Token里
func GenerateToken(keys TokenKeys) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for key, value := range keys {
		claims[key] = value
	}
	//设置过期时间
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	token_str, err := token.SignedString([]byte(secret_key))
	return token_str, err
}

//解析Token里的key
func ParseToken(token string) (TokenKeys, error) {
	var keys TokenKeys = TokenKeys{}
	parseAuth, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})
	if err != nil {
		return keys, errors.New("TOKEN ERROR!")
	}
	claims := parseAuth.Claims.(jwt.MapClaims)
	for key, value := range claims {

		if val, ok := value.(float64); ok {
			value = int(val)
		}

		keys[key] = value
	}
	return keys, nil
}
