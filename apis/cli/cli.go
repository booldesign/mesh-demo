package cli

import (
	"time"

	"github.com/booldesign/protogen/kitex_gen/services/account/usersrv"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var (
	AccountClient usersrv.Client
)

func InitAccountClient() {
	options := localK8sRegisterClient("account-service.banding-cloud.svc.cluster.local:4001")
	c, err := usersrv.NewClient(
		"account-service",
		options...,
	)
	if err != nil {
		panic(err)
	}
	AccountClient = c
}

func localK8sRegisterClient(hostports ...string) []client.Option {
	return []client.Option{
		client.WithSuite(tracing.NewClientSuite()), // tracer
		client.WithRPCTimeout(5 * time.Second),     // rpc timeout
		client.WithConnectTimeout(5 * time.Second), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithHostPorts(hostports...),
	}
}
