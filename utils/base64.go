package utils

import "encoding/base64"

func EncodeBase64(str string) string {
	input := []byte(str)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func DecodeBase64(str string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decodeBytes), nil
}
