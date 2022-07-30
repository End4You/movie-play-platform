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
func GetListLogic(db *gorm.DB, data map[string]interface{}, role uint32) ([]*pb.GetListRsp_Result, error) {
	var list []*pb.GetListRsp_Result
	var selectSQL []string
	// 获取 WHERE 条件
	whereSQL := utils.SpliceWhereSql(data, role)
	// 不同角色 SELECT 字段不同
	if role == uint32(2) {
		selectSQL = config.AdminListFields
	} else {
		selectSQL = config.UserListFields
		whereSQL += fmt.Sprintf(" AND mStatus = %d", 1)
	}
	result := db.Debug().Model(&models.List{}).Select(selectSQL).Where(whereSQL).Scan(&list)

	return list, result.Error
}
