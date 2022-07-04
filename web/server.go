package web

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/logger"
	"time"
)

func Start() {
	hlog.SetLogger(logger.NewHertzFullLogger())
	h := server.Default(server.WithKeepAlive(true), server.WithHostPorts(":9180"))

	configureCors(h)
	registerRoute(h)

	h.Spin()
}

func configureCors(h *server.Hertz) {
	confOrigins := config.Get("cors:origins").([]interface{})
	confMethods := config.Get("cors:methods").([]interface{})

	origins := make([]string, len(confOrigins))
	methods := make([]string, len(confMethods))

	for i := 0; i < len(confOrigins); i++ {
		origins[i] = confOrigins[i].(string)
	}

	for i := 0; i < len(confMethods); i++ {
		methods[i] = confMethods[i].(string)
	}

	h.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	logger.Info("[Web][Server] use cors origins: ", origins)
	logger.Info("[Web][Server] use cors methods: ", methods)
}
