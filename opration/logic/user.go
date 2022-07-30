package logic

import (
	"context"
	"gorm.io/gorm"
	"movie_opration/config"
	"movie_opration/models"
	"movie_opration/utils"
	"time"
)

// RegisterLogic 创建用户，写入到 user 表
func RegisterLogic(db *gorm.DB, userName, password string) error {
	result := db.Debug().Model(&models.User{}).Create(map[string]interface{}{
		"userName":   userName,
		"password":   password,
		"role":       2,
		"createTime": time.Now().Unix(),
		"updateTime": time.Now().Unix(),
	})

	return result.Error
}

// CheckUserNameLogic 检查用户名重复，获取 user 表内是否存在用户名
func CheckUserNameLogic(db *gorm.DB, userName string) (bool, error) {
	var count int64
	result := db.Debug().Model(&models.User{}).Where("userName = ?", userName).Count(&count)

	return count > 0, result.Error
}

// CheckPasswordLogic 检查密码正确性，获取 user 表内是否存在用户名和密码
func CheckPasswordLogic(db *gorm.DB, userName, password string) (bool, error) {
	var dbPassword string
	result := db.Debug().Model(&models.User{}).
		Select("password").
		Where("userName = ?", userName).
		First(&dbPassword)

	return dbPassword == password, result.Error
}

// DeleteTokenLogic 用户登录时，删除旧 token
func DeleteTokenLogic(db *gorm.DB, userName string) error {
	result := db.Debug().Where("userName = ?", userName).Delete(&models.Token{})

	return result.Error
}

// CleanTokenLogic 定时清理，删除过期 token
func CleanTokenLogic(db *gorm.DB) error {
	result := db.Debug().Where("loginTime <= ?", time.Now().Unix()-7*24*60*60).Delete(&models.Token{})

	return result.Error
}

// WriteTokenLogic 用户登录时，写入新 token
func WriteTokenLogic(db *gorm.DB, userName, token string) error {
	result := db.Debug().Model(&models.Token{}).Create(map[string]interface{}{
		"token":     token,
		"userName":  userName,
		"loginTime": time.Now().Unix(),
	})

	return result.Error
}

// PreHandleTokenLogic 检查 http 带来的 token 是否在数据库中
func PreHandleTokenLogic(db *gorm.DB, ctx context.Context) (bool, string, uint32, error) {
	// 获取请求头中的 token
	httpToken, err := utils.GetToken(ctx)
	if err != nil {
		return false, "", 1, config.New(config.ClientNoTokenError)
	}
	// 解密 token 中的 userName 字段
	userName, err := utils.DecodeToken(httpToken, "userName")
	if err != nil {
		return false, "", 1, config.New(config.ClientExtractTokenError)
	}
	// 获取 token 数据表中该 userName 对应的 token
	var token string
	tokenResult := db.Debug().Model(&models.Token{}).
		Select("token").
		Where("userName = ?", userName).
		First(&token)
	// 获取 user 数据表中该 userName 对应的 role
	var role uint32
	roleResult := db.Debug().Model(&models.User{}).
		Select("role").
		Where("userName = ?", userName).
		First(&role)
	// 任意 db 读取错误，则返回错误
	var dbError error
	if tokenResult.Error != nil || roleResult.Error != nil {
		dbError = config.New(config.InnerReadDbError)
	}

	return token == httpToken, userName, role, dbError
}
