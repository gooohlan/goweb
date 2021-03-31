package utils

import (
	"goweb/utils/myerror"
	"goweb/utils/myerror/code"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//常量key
const (
	SIGNED_KEY = "ziudiukeji"
	ISSUER     = "InkDP"
)

// jwt
type jwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	UserId string `json:"user_id"`
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(userId string) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 24).Unix()),
			Issuer:    ISSUER,
		},
		userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(SIGNED_KEY))
	return
}

// 解密
func ParseHSToken(tokenString string) (err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNED_KEY), nil
	})
	if err != nil {
		err = myerror.Logic(code.USER_BUY_NOT_BALANCE)
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = myerror.Logic(code.USER_JWT_ILLEGAL)
		return
	}
	return
}
