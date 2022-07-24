package models

import "time"

type User struct {
	// 用户ID，自增ID
	Uid uint32
	// 用户名
	UserName string
	// 密码
	Password string
	// 角色 1-guest 2-user 3-admin
	Role int8
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
}
