package services

import (
	"context"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.woa.com/crotaliu/pb-hub"
	"movie_opration/config"
	"movie_opration/logic"
	"movie_opration/utils"
)

type UserImpl struct{}

// CreateUser 用户注册
func (s *UserImpl) CreateUser(ctx context.Context, req *pb.CreateUserReq, rsp *pb.CreateUserRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 创建用户前判断用户名是否存在
	result, err := logic.CheckUserNameLogic(db, req.UserName)
	// 用户名重复时返回错误
	if result || err != nil {
		rsp.Code, rsp.Msg = config.InnerOtherError.Code, config.InnerOtherError.Msg
		rsp.Result = 0
		return nil
	}
	// 创建用户逻辑
	err = logic.CreateUserLogic(db, req.UserName, req.Password)
	// 出错处理
	if err != nil {
		log.Errorf("CreateUserLogic 用户注册事件失败：%v", err)
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		rsp.Result = 0

		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	rsp.Result = 1

	return nil
}

// CheckUserName 检查用户名重复
func (s *UserImpl) CheckUserName(ctx context.Context, req *pb.CheckUserNameReq, rsp *pb.CheckUserNameRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 检查用户名重复通用逻辑
	result, err := logic.CheckUserNameLogic(db, req.UserName)
	// 出错处理
	if err != nil {
		log.Errorf("CheckUserName 检查用户名重复事件失败：%v", err)
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		rsp.Result = result
	}
	// 正常返回
	if result {
		rsp.Code, rsp.Msg = config.ClientCheckUserNameError.Code, config.ClientCheckUserNameError.Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	rsp.Result = result

	return nil
}
