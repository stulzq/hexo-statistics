package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stulzq/hexo-statistics/logger"
)

func main() {
	hlog.SetLogger(logger.NewHertzFullLogger())
	h := server.Default(server.WithKeepAlive(true), server.WithHostPorts(":9180"))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		panic("a")
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Spin()
}
