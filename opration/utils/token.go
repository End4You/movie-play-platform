package utils

import (
	"context"
	"encoding/base64"
	"git.code.oa.com/trpc-go/trpc-go/http"
	"github.com/gofrs/uuid"
	"movie_opration/config"
	"strconv"
	"strings"
	"time"
)

// GetToken 获取请求用户 token
func GetToken(ctx context.Context) (string, error) {
	request := http.Request(ctx)
	token := request.Header["X-Token"]
	if token != nil {
		return token[0], nil
	}

	return "", config.New(config.ClientNoTokenError)
}

// GenerateToken 生成 token
func GenerateToken(userName string) string {
	uid, _ := uuid.NewV4()                                   // 生成 UUID
	timestampNow := strconv.FormatInt(time.Now().Unix(), 10) // 时间戳文本

	return base64.StdEncoding.EncodeToString(
		[]byte("userName=" + userName + ";" +
			"loginTime=" + timestampNow + ";" +
			"uuid=" + uid.String()))
}

// DecodeToken 解密 token 字段信息
func DecodeToken(token string, key string) (string, error) {
	// 解密 token
	decodeStr, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", config.New(config.ClientExtractTokenError)
	}
	// 解密 token 内容转为文本
	tokenStr := string(decodeStr)
	// token map
	tokenMap := make(map[string]string)
	// 判断 token 格式是否正确
	if strings.Contains(tokenStr, "userName") &&
		strings.Contains(tokenStr, "loginTime") &&
		strings.Contains(tokenStr, "uuid") &&
		strings.Contains(tokenStr, ";") &&
		strings.Contains(tokenStr, "=") {
		tokenSplit := strings.Split(tokenStr, ";")
		for _, tokenItem := range tokenSplit {
			tokenItemSplice := strings.Split(tokenItem, "=")
			if len(tokenItemSplice) != 2 {
				return "", config.New(config.ClientExtractTokenError)
			}
			tokenMap[tokenItemSplice[0]] = tokenItemSplice[1]
		}
	}

	return tokenMap[key], nil
}
