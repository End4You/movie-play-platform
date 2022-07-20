package tests

import (
	"context"
	"movie_opration/services"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"

	pb "git.woa.com/crotaliu/pb-hub"
)

//go:generate go mod tidy

//go:generate mockgen -destination=stub/git.woa.com/crotaliu/pb-hub/movie_mock.go -package=pb_hub -self_package=git.woa.com/crotaliu/pb-hub --source=stub/git.woa.com/crotaliu/pb-hub/movie.trpc.go

func Test_movieImpl_GetMovieList(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	movieService := pb.NewMockMovieService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := movieService.EXPECT().GetMovieList(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.GetMovieListReq, rsp *pb.GetMovieListRsp) error {
		s := &services.MovieImpl{}
		return s.GetMovieList(ctx, req, rsp)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.GetMovieListReq
		rsp *pb.GetMovieListRsp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp := &pb.GetMovieListRsp{}
			if err := movieService.GetMovieList(tt.args.ctx, tt.args.req, rsp); (err != nil) != tt.wantErr {
				t.Errorf("movieImpl.GetMovieList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("movieImpl.GetMovieList() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
