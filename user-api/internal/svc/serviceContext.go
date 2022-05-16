package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/model"
	"zero-demo/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel

	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
