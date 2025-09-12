package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId uint `json:"userid"`
	jwt.RegisteredClaims
}

var signkey = []byte("golang")

func GenerateJWT(userId uint) (string, error) {
	//创建JWT的声明
	claims := UserClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//签名
	//密钥需要是[]byte类型
	if ss, err := token.SignedString(signkey); err != nil {
		return "", err
	} else {
		return "Bearer " + ss, nil
	}

}
