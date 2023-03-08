package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/booldesign/protogen/kitex_gen/base"
	"github.com/booldesign/protogen/kitex_gen/services/account"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/baggage"
)

// UserGetBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserGetBe(ctx context.Context, req *account.UserGetBeReq) (resp *account.UserGetBeResp, err error) {
	value, ok := metainfo.GetPersistentValue(ctx, "branch")
	fmt.Printf("mate-branch:%+v, %t\n", value, ok)

	value, ok = metainfo.GetPersistentValue(ctx, "baggage")
	fmt.Printf("mate-baggage:%+v, %t\n", value, ok)

	klog.CtxDebugf(ctx, "baggage: %s", baggage.FromContext(ctx).String())

	resp = &account.UserGetBeResp{BaseResp: &base.BaseResp{StatusCode: 0, Message: "SUCCESS"}, Data: &account.UserBe{}}
	resp.Data.Id = 0
	resp.Data.Name = "edison rpc " + os.Getenv("OTEL_RESOURCE_ATTRIBUTES")

	return
}

// UserListBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserListBe(ctx context.Context, req *account.UserListBeReq) (resp *account.UserListBeResp, err error) {
	// TODO: Your code here...
	return
}
