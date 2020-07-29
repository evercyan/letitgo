package crypto

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Key    string
	Expire int
}

func NewJWT(key string, expire int) *JWT {
	return &JWT{
		Key:    key,
		Expire: expire,
	}
}

/**
 * payload: map[string]interface{}{
 *     "iss": "jwt",         // 签发者
 *     "iat": 1595838971,    // 签发时间
 *     "exp": 1595838972,    // 过期时间
 *     "nbf": 1595838972,    // 校验时间, 该时间前此 token 无效
 *     "sub": "www.xxx.com", // 面向用户
 *     "jti": "xxxx",        // 该 token 唯一标识
 *     "xxx": "xxx",         // 可附加信息
 * }
 *
 */
func (t *JWT) GetToken(payload map[string]interface{}) (string, error) {
	payload["iat"] = time.Now().Unix()
	payload["exp"] = time.Now().Add(time.Duration(t.Expire) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	tokenString, err := token.SignedString([]byte(t.Key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *JWT) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.Key), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
