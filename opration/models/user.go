package models

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
	CreateTime int64
	// 更新时间
	UpdateTime int64
}

type Token struct {
	// 用户token
	Token string
	// 用户名
	UserName string
	// 登录时间
	LoginTime int64
}
