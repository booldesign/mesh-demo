package handlers

import (
	"context"

	"github.com/booldesign/protogen/kitex_gen/base"
	"github.com/booldesign/protogen/kitex_gen/services/account"
)

// UserGetBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserGetBe(ctx context.Context, req *account.UserGetBeReq) (resp *account.UserGetBeResp, err error) {
	resp = &account.UserGetBeResp{BaseResp: &base.BaseResp{StatusCode: 0, Message: "SUCCESS"}, Data: &account.UserBe{}}
	resp.Data.Id = 1
	resp.Data.Name = "edison1"

	return
}

// UserListBe implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserListBe(ctx context.Context, req *account.UserListBeReq) (resp *account.UserListBeResp, err error) {
	// TODO: Your code here...
	return
}
