package main

import (
	"fmt"
	"net"

	"example.com/services/account/handlers"
	trace "github.com/kitex-contrib/tracer-opentracing"

	account "github.com/booldesign/protogen/kitex_gen/services/account/usersrv"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func main() {
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
		server.WithLimit(&limit.Option{MaxConnections: 100, MaxQPS: 100}),
		server.WithSuite(trace.NewDefaultServerSuite()), // tracer
		//server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}
}
