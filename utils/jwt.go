package utils

import (
	"errors"
	"time"

	"golang-web/config"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明结构
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int, username string, cfg *config.Config) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(time.Duration(cfg.JWT.Expire) * time.Hour)

	// 创建声明
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "golang-web",
			Subject:   username,
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString([]byte(cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string, cfg *config.Config) (*Claims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(cfg.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌
	if !token.Valid {
		return nil, errors.New("无效的令牌")
	}

	// 提取声明
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("无法提取声明")
	}

	return claims, nil
}

// RefreshToken 刷新JWT令牌
func RefreshToken(tokenString string, cfg *config.Config) (string, error) {
	// 验证当前令牌
	claims, err := ValidateToken(tokenString, cfg)
	if err != nil {
		return "", err
	}

	// 生成新令牌
	return GenerateToken(claims.UserID, claims.Username, cfg)
}
