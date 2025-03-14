package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT（JSON Web Token）
// JWT 是一种用于 身份认证 和 信息安全传输 的 Token，常用于：
//	•	用户认证（Auth）
//	•	API 访问控制
//	•	分布式系统授权（微服务）
//
// JWT 由 三部分 组成，每部分用 . 分隔：Header.Payload.Signature
// 示例：
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsImV4cCI6MTY5MDAwMDAwMH0.H4w1h5YBvGZrJDhfCp2QUgUfdU9Q-WPlu2PEoKq1Iw8
// 	•	Header（头部）：包含 Token 类型（JWT） 和 签名算法（HS256、RS256）。
//	•	Payload（负载）：存储用户信息，如 user_id、exp（过期时间）。
//	•	Signature（签名）：用 密钥 生成的哈希值，防止 Token 被篡改。
//
// 安装： go get github.com/golang-jwt/jwt/v4

// 生成jwt
// 定义密钥（生产环境请存放在安全的地方）
var secretKey = []byte("my_secret_key")

// 生成 JWT Token
func GenerateToken(userID int) (string, error) {
	// 创建 JWT 负载
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // 2 小时后过期
		"iat":     time.Now().Unix(),                    // 签发时间
		"iss":     "my_app",                             // 签发者
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥进行签名
	return token.SignedString(secretKey)
}

func main() {
	token, err := GenerateToken(123)
	if err != nil {
		fmt.Println("生成 Token 失败:", err)
		return
	}
	fmt.Println("JWT Token:", token)
}
