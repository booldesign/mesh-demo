package handlers

import (
	"context"
	"fmt"

	"github.com/booldesign/protogen/kitex_gen/base"
	"github.com/booldesign/protogen/kitex_gen/services/account"
	"github.com/bytedance/gopkg/cloud/metainfo"
)

// UserGetBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserGetBe(ctx context.Context, req *account.UserGetBeReq) (resp *account.UserGetBeResp, err error) {
	value, ok := metainfo.GetPersistentValue(ctx, "branch")
	fmt.Printf("mate:%+v, %t\n", value, ok)

	resp = &account.UserGetBeResp{BaseResp: &base.BaseResp{StatusCode: 0, Message: "SUCCESS"}, Data: &account.UserBe{}}
	resp.Data.Id = 2
	resp.Data.Name = "edison2"

	return
}

// UserListBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserListBe(ctx context.Context, req *account.UserListBeReq) (resp *account.UserListBeResp, err error) {
	// TODO: Your code here...
	return
}
