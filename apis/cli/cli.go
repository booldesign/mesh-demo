package cli

import (
	"example.com/apis/usercenter/pkg/metadata"
	"github.com/booldesign/protogen/kitex_gen/services/account/usersrv"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	xds2 "github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
)

var (
	AccountClient usersrv.Client
)

func InitAccountClient() {
	options := localK8sRegisterClient()
	c, err := usersrv.NewClient(
		"account-service.banding-cloud.svc.cluster.local:4001",
		options...,
	)

	if err != nil {
		panic(err)
	}
	AccountClient = c
	return
}

func localK8sRegisterClient() []client.Option {
	// initialize xds module
	if err := xds.Init(xds.WithXDSServerAddress("istiod.istio-system.svc:15010")); err != nil {
		klog.Fatal(err)
	}

	return []client.Option{
		client.WithSuite(tracing.NewClientSuite()), // tracer
		client.WithXDSSuite(xds2.ClientSuite{
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(metadata.ExtractFromPropagator),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
		// client.WithTag("branch", "v2"),
	}
}
