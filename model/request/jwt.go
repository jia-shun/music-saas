package request

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID         uint
	Username   string
	BufferTime int64
	jwt.StandardClaims
}
