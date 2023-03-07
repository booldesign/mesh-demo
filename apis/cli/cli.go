package cli

import (
	"time"

	"github.com/booldesign/protogen/kitex_gen/services/account/usersrv"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/metadata"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	xds2 "github.com/cloudwego/kitex/pkg/xds"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
)

var (
	AccountClient usersrv.Client
)

func InitAccountClient() usersrv.Client {
	options := localK8sRegisterClient("account-service:4001")
	c, err := usersrv.NewClient(
		"account-rpc",
		options...,
	)
	if err != nil {
		panic(err)
	}
	return c
}

func localK8sRegisterClient(hostports ...string) []client.Option {
	// initialize xds module
	if err := xds.Init(); err != nil {
		klog.Fatal(err)
	}

	return []client.Option{
		client.WithRPCTimeout(5 * time.Second),     // rpc timeout
		client.WithConnectTimeout(5 * time.Second), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithHostPorts(hostports...),
		//client.WithTransportProtocol(transport.GRPC),
		//client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithXDSSuite(xds2.ClientSuite{
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(metadata.ExtractFromPropagator),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
		// client.WithTag("branch", "v2"),
	}
}
