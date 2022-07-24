package services

import (
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type MovieImpl struct{}

// GetMovieList 获取视频列表
func (s *MovieImpl) GetMovieList(ctx context.Context, req *pb.GetMovieListReq, rsp *pb.GetMovieListRsp) error {
	// implement business logic here ...
	// ...
	return nil
}
