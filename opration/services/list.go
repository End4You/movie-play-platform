package services

import (
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type ListImpl struct{}

// GetList 获取视频列表
func (s *ListImpl) GetList(ctx context.Context, req *pb.GetListReq, rsp *pb.GetListRsp) error {
	// implement business logic here ...
	// ...
	return nil
}
