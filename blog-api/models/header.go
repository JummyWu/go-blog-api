package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
)

/*
SetToken 生成jwt
*/
func SetToken(uid string) (string, error) {
	// red := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "wujianheng",
	// 	DB:       0,
	// })
	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, error := newJwt.SignedString([]byte("wujianheng"))
	// err := red.Set(uid, token, 60*60*time.Minute).Err()
	// if err != nil {
	// 	logs.Info(err)
	// }
	if error != nil {
		logs.Info("生成错误")
	}
	return token, error
}

/*
ParseToken 检查请求头是否带token
*/
func ParseToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["uid"])
		}
		hmac := []byte("wujianheng")
		return hmac, nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

/*
GetUID 从token中获取用户的UID
*/
func GetUID(tokenStr string) string {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["uid"])
		}
		hmac := []byte("wujianheng")
		return hmac, nil
	})
	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := claims["uid"]
		return uid.(string)
	} else {
		return ""
	}
}
