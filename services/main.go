package main

import (
	"context"
	"fmt"
	"net"

	"example.com/services/account/handlers"
	account "github.com/booldesign/protogen/kitex_gen/services/account/usersrv"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("account-service"),
		provider.WithInsecure(),
	)
	defer func(p provider.OtelProvider, ctx context.Context) {
		_ = p.Shutdown(ctx)
	}(p, ctx)

	options := localOption()
	svr := account.NewServer(new(handlers.UserSrvImpl), options...)
	err := svr.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func localOption() []server.Option {
	address, err := net.ResolveTCPAddr("tcp", ":4001")
	if err != nil {
		panic(err)
	}
	return []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "account-service"}),
		server.WithServiceAddr(address),
		server.WithSuite(tracing.NewServerSuite()),
	}
}
