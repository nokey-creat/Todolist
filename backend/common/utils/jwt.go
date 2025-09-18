package utils

import (
	"errors"
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
		//return "Bearer " + ss, nil //前端会加bearer
		return ss, nil
	}

}

func ParseJWT(tokenString string) (uint, error) {
	//解析并验证
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		//该函数接收已解析但未验证的 Token
		//这里限制签名算法
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("JWT with Unexpected signed method")
		}
		return signkey, nil
	})
	if err != nil {
		return 0, err
	} else if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("unknown claims type, cannot proceed")
	}

}
