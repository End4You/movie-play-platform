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
	result, count, err := logic.GetListLogic(db, mapData, role, req.PageNo, req.PageSize)
	if err != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回列表信息，前端正常展示列表，但清空用户信息
	if !tokenBol || err != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, config.ClientUserInfoError.Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	rsp.Result = &pb.GetListRsp_Result{
		List:  result,
		Count: count,
	}

	return nil
}
