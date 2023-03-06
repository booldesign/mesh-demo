package handlers

import (
	"context"

	"example.com/apis/usercenter/cli"
	"github.com/booldesign/protogen/kitex_gen/services/account"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserGetReq struct {
	Id uint64 `json:"id"`
}

type UserGetResp struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func UserGet(ctx context.Context, c *app.RequestContext) {
	req := UserGetReq{}
	if err := c.Bind(&req); err != nil {
		c.JSON(400, "params err"+err.Error())
		return
	}

	userCenterResp, err := cli.AccountClient.UserGetBe(ctx, &account.UserGetBeReq{Id: req.Id})
	if err != nil {
		c.JSON(500, "userCenterResp err")
		return
	}
	var data *UserGetResp
	if userCenterResp.GetData() != nil {
		data = &UserGetResp{Id: userCenterResp.GetData().GetId(), Name: userCenterResp.GetData().GetName() + "/2"}
	}
	c.JSON(200, data)
	return
}
