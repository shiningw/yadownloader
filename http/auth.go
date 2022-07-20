package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/filebrowser/filebrowser/v2/users"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

const (
	TokenExpirationTime = time.Hour * 2
)

type userInfo struct {
	ID           uint              `json:"id"`
	Locale       string            `json:"locale"`
	ViewMode     users.ViewMode    `json:"viewMode"`
	SingleClick  bool              `json:"singleClick"`
	Perm         users.Permissions `json:"perm"`
	Commands     []string          `json:"commands"`
	LockPassword bool              `json:"lockPassword"`
	HideDotfiles bool              `json:"hideDotfiles"`
	DateFormat   bool              `json:"dateFormat"`
}

type authToken struct {
	User userInfo `json:"user"`
	jwt.RegisteredClaims
}

type extractor []string

func (e extractor) ExtractToken(r *http.Request) (string, error) {
	token, _ := request.HeaderExtractor{"X-Auth"}.ExtractToken(r)

	// Checks if the token isn't empty and if it contains two dots.
	// The former prevents incompatibility with URLs that previously
	// used basic auth.
	if token != "" && strings.Count(token, ".") == 2 {
		return token, nil
	}

	auth := r.URL.Query().Get("auth")
	if auth != "" && strings.Count(auth, ".") == 2 {
		return auth, nil
	}

	cookie, _ := r.Cookie("auth")
	if cookie != nil && strings.Count(cookie.Value, ".") == 2 {
		return cookie.Value, nil
	}

	return "", request.ErrNoTokenInRequest
}

type authentication struct {
	tlsKey []byte
	r      *http.Request
}

func (c authentication) auth() bool {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return c.tlsKey, nil
	}
	//cookie, _ := c.r.Cookie("auth")
	//log.Println(cookie.String())
	var tk authToken
	token, err := request.ParseFromRequest(c.r, &extractor{}, keyFunc, request.WithClaims(&tk))

	if err != nil || !token.Valid {
		return false
	}
	return true
}
