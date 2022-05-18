package rest

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	// HeaderContentType is the standard content type header
	HeaderContentType = `Content-Type`
)

var reBearToken = regexp.MustCompile("^[Bb]earer (.+)$")

func ParseBearerToken(r *http.Request) (string, error) {
	auth := r.Header.Get("authorization")
	if !reBearToken.MatchString(auth) {
		return "", fmt.Errorf("authorization required")
	}
	token := reBearToken.FindAllStringSubmatch(auth, -1)[0][1]
	if strings.TrimSpace(token) == "" {
		return "", fmt.Errorf("token can not be blank")
	}
	return token, nil
}
