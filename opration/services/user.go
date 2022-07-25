package services

import (
	"context"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.woa.com/crotaliu/pb-hub"
	"movie_opration/config"
	"movie_opration/logic"
	"movie_opration/utils"
	"time"
)

type UserImpl struct{}

// Register 用户注册
func (s *UserImpl) Register(ctx context.Context, req *pb.RegisterReq, rsp *pb.RegisterRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 创建用户前判断用户名是否存在
	result, err := logic.CheckUserNameLogic(db, req.UserName)
	// 用户名重复时返回错误
	if result || err != nil {
		rsp.Code, rsp.Msg = config.InnerOtherError.Code, config.InnerOtherError.Msg
		rsp.Result = false
		return nil
	}
	// 创建用户逻辑
	err = logic.RegisterLogic(db, req.UserName, req.Password)
	// 出错处理
	if err != nil {
		log.Errorf("CreateUserLogic 用户注册事件失败：%v", err)
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		rsp.Result = false

		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	rsp.Result = true

	return nil
}

// CheckUserName 检查用户名重复
func (s *UserImpl) CheckUserName(ctx context.Context, req *pb.CheckUserNameReq, rsp *pb.CheckUserNameRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 检查用户名重复逻辑
	result, err := logic.CheckUserNameLogic(db, req.UserName)
	// 出错处理
	if err != nil {
		log.Errorf("CheckUserNameLogic 检查用户名重复事件失败：%v", err)
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

// Login 用户登录
func (s *UserImpl) Login(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 检查密码正确性
	result, err := logic.CheckPasswordLogic(db, req.UserName, req.Password)
	if err != nil {
		log.Errorf("CheckPasswordLogic 检查密码正确性事件失败：%v", err)
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
	}
	if result { // 密码正确
		// 生成 token
		token := utils.GenerateToken(req.UserName)
		// 删除旧 token
		err = logic.DeleteTokenLogic(db, req.UserName)
		if err != nil {
			log.Errorf("DeleteTokenLogic 删除旧 token 事件失败：%v", err)
			rsp.Code, rsp.Msg = config.InnerDeleteDbError.Code, config.InnerDeleteDbError.Msg
		}
		// 写入新 token
		err = logic.WriteTokenLogic(db, req.UserName, token)
		if err != nil {
			log.Errorf("WriteTokenLogic 写入新 token 事件失败：%v", err)
			rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		}
		// 正常返回
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
		rsp.Result = &pb.LoginRsp_Result{
			Token:     token,
			LoginTime: time.Now().Unix(),
		}
	} else { // 密码错误
		rsp.Code, rsp.Msg = config.ClientLoginError.Code, config.ClientLoginError.Msg
	}

	return nil
}
