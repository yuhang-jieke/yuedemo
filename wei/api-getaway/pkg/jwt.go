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
	// 颁发一个有限期一小时的证书
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":    time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	return tokenString, err
}
func PersonToken(tokenString string) (jwt.MapClaims, error) {
	//jwt.Parse()
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
func CreateToken(tokenString string) (string, error) {
	//jwt.Parse()
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token格式错误")
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("无法获取用户ID")
	}
	return TokenHandler(userId)
}
func RefreshToken(oldToken string) (string, error) {
	// 解析旧token时，忽略过期错误，仅验证签名（用生成时的APP_KEY，和TokenHandler保持一致）
	token, _ := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil // 注意：这里用APP_KEY，和生成时保持一致
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token格式错误")
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("无法获取用户ID")
	}

	// 复用你现有的TokenHandler生成新的Token
	return TokenHandler(userId)
}
