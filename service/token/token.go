package token

import (
	"simple-core/service/errmsg"
	"simple-core/setting"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserClaims struct {
	UserId    int64
	UserEmail string
	UserRole  int
	jwt.StandardClaims
}

func Sign(uid int64, email string, role int) (string, error) {

	expAt := time.Now().Add(time.Duration(12) * time.Hour).Unix()

	claims := UserClaims{
		UserId:    uid,
		UserEmail: email,
		UserRole:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "simple web token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(setting.SecureKey))
}

func checkToken(tokenStr string) (*UserClaims, error) {
	if tokenStr == "" {
		return nil, errmsg.TokenWrongError
	}
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(ts *jwt.Token) (interface{}, error) {
		return []byte(setting.SecureKey), nil
	})
	if err != nil {
		return nil, errmsg.TokenParseWrongError
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errmsg.TokenParseWrongError
}

func ParseToken(c *gin.Context) *UserClaims {
	tokenHeader := c.Request.Header.Get("Authorization")
	cToken := strings.SplitN(tokenHeader, " ", 2)
	if len(cToken) != 2 && cToken[0] != "Bearer" {
		return nil
	}
	u, err := checkToken(cToken[1])
	if err != nil {
		return nil
	}

	if time.Now().Unix() > u.ExpiresAt {
		return nil
	}

	return u
}
