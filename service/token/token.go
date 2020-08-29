package token

import (
	"simple-core/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId    int64
	UserEmail string
	jwt.StandardClaims
}

func Sign(uid int64, email string) (string, error) {

	expAt := time.Now().Add(time.Duration(12) * time.Hour).Unix()

	claims := UserClaims{
		UserId:    uid,
		UserEmail: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "simple web token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(setting.SecureKey))
}
