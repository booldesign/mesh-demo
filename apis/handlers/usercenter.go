package handlers

import (
	"context"
	"fmt"
	"os"

	"example.com/apis/usercenter/cli"

	"github.com/booldesign/protogen/kitex_gen/services/account"
	"github.com/bytedance/gopkg/cloud/metainfo"
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
	if c.Request.Header.Get("branch") != "" {
		fmt.Printf("branch:%s\n", c.Request.Header.Get("branch"))
		ctx = metainfo.WithPersistentValue(ctx, "branch", c.Request.Header.Get("branch"))
	}

	userCenterResp, err := cli.InitAccountClient().UserGetBe(ctx, &account.UserGetBeReq{Id: req.Id})
	if err != nil {
		c.JSON(500, "userCenterResp err"+err.Error())
		return
	}
	var data *UserGetResp
	if userCenterResp.GetData() != nil {
		data = &UserGetResp{Id: userCenterResp.GetData().GetId(), Name: userCenterResp.GetData().GetName() + "/api " + os.Getenv("OTEL_RESOURCE_ATTRIBUTES")}
	}
	c.JSON(200, data)
	return
}
