package services

import (
	"context"
	"encoding/json"
	pb "git.woa.com/crotaliu/pb-hub"
	"movie_opration/config"
	"movie_opration/logic"
	"movie_opration/utils"
)

type ListImpl struct{}

// GetList 获取视频列表
func (s *ListImpl) GetList(ctx context.Context, req *pb.GetListReq, rsp *pb.GetListRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, err := logic.PreHandleTokenLogic(db, ctx)
	if !tokenBol || err != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, config.ClientUserInfoError.Msg
		return nil
	}
	// struct 转 map
	mapData := make(map[string]interface{})
	reqData, marshalErr := json.Marshal(&req)
	if marshalErr != nil {
		rsp.Code, rsp.Msg = config.InnerMarshalError.Code, config.InnerMarshalError.Msg
		return nil
	}
	unMarshalErr := json.Unmarshal(reqData, &mapData)
	if unMarshalErr != nil {
		rsp.Code, rsp.Msg = config.InnerUnmarshalError.Code, config.InnerUnmarshalError.Msg
		return nil
	}
	// 查询电影列表逻辑
	result, err := logic.GetListLogic(db, mapData, role)
	if err != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	rsp.Result = result

	return nil
}
