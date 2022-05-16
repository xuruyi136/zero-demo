package logic

import (
	"context"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	// post 127.0.0.1:8888/user/info    body  {  "userId":1 }
	/* m := map[int64]string{
		1: "张三",
		2: "李四",
	}
	nikname := "unknow"
	if name, ok := m[req.UserId]; ok {
		nikname = name
	}
	return &types.UserInfoResponse{
		UserId:   req.UserId,
		Nickname: nikname,
	}, nil */
	/* user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInfoResponse{
		UserId:   user.Id,
		Nickname: user.Nickname,
	}, nil */

	//调用rpc
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		UserId:   userResp.Id,
		Nickname: userResp.Nickname,
	}, nil
}
