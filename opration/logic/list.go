package logic

import (
	"fmt"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
	"movie_opration/config"
	"movie_opration/models"
	"movie_opration/utils"
)

// GetListLogic 获取电影列表
func GetListLogic(db *gorm.DB, data map[string]interface{}, role uint32, pageNo, pageSize int32) ([]*pb.GetListRsp_List, int64, error) {
	var list []*pb.GetListRsp_List
	var count int64
	var selectSQL []string
	// 获取 WHERE 条件
	whereSQL := utils.SpliceWhereSql(data, role)
	// 不同角色 SELECT 字段不同
	if role == uint32(3) {
		selectSQL = config.AdminListFields
	} else {
		selectSQL = config.UserListFields
		whereSQL += fmt.Sprintf(" AND mStatus = %d", 1)
	}
	// 分页列表
	listResult := db.Debug().Model(&models.List{}).Select(selectSQL).Where(whereSQL).
		Offset(int((pageNo - 1) * pageSize)).Limit(int(pageSize)).Scan(&list)
	// 总量
	countResult := db.Debug().Model(&models.List{}).Select(selectSQL).Where(whereSQL).Count(&count)
	// 任意 db 读取错误，则返回错误
	var dbError error
	if listResult.Error != nil || countResult.Error != nil {
		dbError = config.New(config.InnerReadDbError)
	}
	return list, count, dbError
}
