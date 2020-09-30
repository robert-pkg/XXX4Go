package app

import (
	"github.com/pkg/errors"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/service"
	"github.com/robert-pkg/micro-go/registry"
	grpc_server "github.com/robert-pkg/micro-go/rpc/server/grpc"

	pb "github.com/robert-pkg/XXX4Go/interface/XXXUserServer/api"
)

// Start 启动服务器
func (app *app) startGrpcServer(r registry.Registry, serviceName string, svc *service.Service) *grpc_server.Server {

	grpcSvr := grpc_server.NewServer(r)
	pb.RegisterXXXUserServerServer(grpcSvr.Server(), svc)

	if err := grpcSvr.Start(serviceName); err != nil {
		panic(errors.Wrap(err, "grpc server start fail"))
	}

	return grpcSvr
}
