package utils

import (
	"fmt"
	tgorm "git.code.oa.com/trpc-go/trpc-database/gorm"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"movie_opration/config"
	"strings"
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

// SpliceWhereSql 拼接 SQL 查询语句
func SpliceWhereSql(data map[string]interface{}, role uint32) string {
	var whereSql string          // 筛选条件语句
	var whereConditions []string // 计算筛选条件需要的字段
	// 数据为切片时，使用 OR 筛选，数据不为切片时，使用相等筛选
	for key := range data {
		item, ok := data[key].([]interface{})
		if ok && len(item) != 0 {
			var strDBSlice []string
			var strDB string
			for _, value := range item {
				strDBSlice = append(strDBSlice, fmt.Sprintf("%s = %v", key, value))
			}
			if len(item) == 1 {
				strDB = strings.Join(strDBSlice, " OR ")
			} else {
				strDB = `(` + strings.Join(strDBSlice, " OR ") + `)`
			}
			whereConditions = append(whereConditions, strDB)
		} else {
			switch key {
			case "pageNo", "pageSize":
				continue
			case "mName":
				str, result := data[key].(string)
				if result && str != "" {
					whereConditions = append(whereConditions, fmt.Sprintf("%s='%s'", key, str))
				}
			case "mDoubanScore":
				num, result := data[key].(float64)
				if result && num != 0 {
					whereConditions = append(whereConditions, fmt.Sprintf("%s>='%v'", key, uint32(num)))
				}
			default:
				num, result := data[key].(float64)
				if result && num != 0 {
					whereConditions = append(whereConditions, fmt.Sprintf("%s='%v'", key, uint32(num)))
				}
			}
		}
	}
	// 组成筛选语句
	if len(whereConditions) != 0 {
		whereSql += strings.Join(whereConditions, " AND ")
	}

	return whereSql
}
