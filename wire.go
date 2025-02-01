//go:build wireinject
// +build wireinject

package main

import (
	"github.com/asynccnu/be-counter/grpc"
	"github.com/asynccnu/be-counter/ioc"
	"github.com/asynccnu/be-counter/pkg/grpcx"
	"github.com/asynccnu/be-counter/repository/cache"
	"github.com/asynccnu/be-counter/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewCounterServiceServer,
		service.NewCachedCounterService,
		cache.NewRedisCounterCache,
		// 第三方
		ioc.InitRedis,
		ioc.InitLogger,
		ioc.InitEtcdClient,
	)
	return grpcx.Server(nil)
}
