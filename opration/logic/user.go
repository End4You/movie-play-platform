package logic

import (
	"gorm.io/gorm"
	"movie_opration/models"
	"time"
)

// CreateUserLogic 创建用户，写入到 user 表
func CreateUserLogic(db *gorm.DB, userName, password string) error {
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
