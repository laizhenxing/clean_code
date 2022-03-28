package identity

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	ccErr "cc/application/error"
)

const (
	ExpiresTime time.Duration = time.Hour * 24 * 7
)

type Claims struct {
	jwt.StandardClaims
	UserID    uint64      `json:"user_id"`
	Authrized bool        `json:"authrized"`
	Data      interface{} `json:"data"`
}

func NewClaims(userID uint64, authrized bool, data interface{}) *Claims {
	return &Claims{UserID: userID, Authrized: authrized, Data: data}
}

func GenToken(claims *Claims, key []byte) (string, error) {
	if claims.ExpiresAt == 0 {
		// 设置过期时间
		expire := time.Now().Add(ExpiresTime).Unix()

		claims.IssuedAt = time.Now().Unix()
		claims.ExpiresAt = expire
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token and return
	return token.SignedString(key)
}

func ParseToken(tokenStr string, key []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Refresh 刷新token
func Refresh(tokenStr string, key []byte) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", err
	}
	if err := token.Claims.Valid(); err != nil {
		return "", err
	}
	// 更新过期时间
	claims.ExpiresAt = time.Now().Add(ExpiresTime).Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newToken.SignedString(key)
}

// VerifyToken 验证token的有效性
func ValidToken(tokenStr string, key []byte) bool {
	err := verify(tokenStr, key)
	return err == nil
}

func ValidTokenError(tokenStr string, key []byte) error {
	return verify(tokenStr, key)
}

func verify(tokenStr string, key []byte) error {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*Claims); !ok {
		return ccErr.Err.E(ccErr.ErrParseClaims)
	}

	if err = token.Claims.Valid(); err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// 错误的token
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return ccErr.Err.E(ccErr.ErrErrorToken)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// token 过期
				return ccErr.Err.E(ccErr.ErrExpiredToken)
			} else {
				// 无法处理的token
				return ccErr.Err.E(ccErr.ErrUnknownToken)
			}
		}
	}

	return nil
}
