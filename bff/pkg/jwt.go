package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	APP_KEY = "www.topgoer.com"
)

// TokenHandler是我们获取用户名和密码的处理程序，如果有效，则返回用于将来请求的令牌。
func TokenHandler(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":    time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	return tokenString, err
}
func GenToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return nil, errors.New("登录异常")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return nil, errors.New("登录超时")
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return nil, errors.New("登录异常")
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, errors.New("登录异常")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, nil
	} else {
		fmt.Println(err)
	}

	return nil, nil
}

func ShuaXin(tokenString string) (string, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token格式错误")
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("无法获取ID")
	}
	return TokenHandler(userId)
}
