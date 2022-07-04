package web

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stulzq/hexo-statistics/logger"
)

func Start() {
	hlog.SetLogger(logger.NewHertzFullLogger())
	h := server.Default(server.WithKeepAlive(true), server.WithHostPorts(":9180"))

	registerRoute(h)

	h.Spin()
}
