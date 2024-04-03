package svc

import (
	"connection/internal/config"
	innerGorm "connection/internal/gorm"
	"connection/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	// 数据库连接
	DB                       *gorm.DB
	AuthInterceptor          rest.Middleware
	RequestHeaderInterceptor rest.Middleware
	CheckGrpcInterceptor     rest.Middleware
	// ChainServiceClient 链服务客户端
	//ChainServiceClient chainCli.Chain
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 创建数据库连接
	db, err := innerGorm.GetConnection(c.MysqlDNS)
	if err != nil {
		panic(err)
	}

	svcCtx := &ServiceContext{
		Config:                   c,
		DB:                       db,
		AuthInterceptor:          middleware.NewAuthInterceptorMiddleware().Handle,
		RequestHeaderInterceptor: middleware.NewRequestHeaderInterceptorMiddleware().Handle,
	}

	return svcCtx
}
