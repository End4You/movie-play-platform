package utils

import (
	tgorm "git.code.oa.com/trpc-go/trpc-database/gorm"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"movie_opration/config"
)

// ConnDB 建立 gorm DB 连接
func ConnDB() *gorm.DB {
	mysqlConf, ok := client.DefaultClientConfig()[config.MySQLBasic]
	if !ok {
		panic("missing mysql config")
	}
	mysqlDsn := mysqlConf.Target[6:]

	// 使用连接池，logger
	connPool := tgorm.NewConnPool(mysqlDsn)
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: connPool,
			}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
			Logger: tgorm.DefaultTRPCLogger,
		},
	)

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
