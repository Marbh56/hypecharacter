package Auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

func GetEmailAndPassword(headers http.Header) (string, string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "Basic" {
		return "", "", errors.New("authorization header is malformed")
	}
	decoded, err := base64.StdEncoding.DecodeString(vals[1])
	if err != nil {
		return "", "", errors.New("authorization header credential decoding error")
	}

	credParts := strings.Split(string(decoded), ":")
	if len(credParts) != 2 {
		return "", "", errors.New("authorization header credential format error")
	}

	email := credParts[0]
	password := credParts[1]

	return email, password, nil

}
