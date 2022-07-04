package web

import (
	"context"
	"github.com/stulzq/hexo-statistics/util"
	"github.com/stulzq/hexo-statistics/web/handlers/statistics"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func registerRoute(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "pong")
	})

	h.GET("/h", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, util.BytesToStr(ctx.Request.Header.RawHeaders()))
	})

	h.GET("/stat/counter", statistics.Counter)
	h.GET("/stat/get", statistics.Get)
}
