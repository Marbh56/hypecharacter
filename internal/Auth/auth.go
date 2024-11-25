package Auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API Key from headers

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Authorization header format error")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Authorization header format error")
	}
	return vals[1], nil
}
