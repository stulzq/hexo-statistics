package statistics

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/logger"
	"github.com/stulzq/hexo-statistics/util"
	"github.com/stulzq/hexo-statistics/web/handlers/statistics/model"
	"net/url"
)

var allowSite map[string]int

func init() {
	data := config.Get("statistics:site").([]interface{})
	allowSite = make(map[string]int, len(data))

	for i := 0; i < len(data); i++ {
		allowSite[data[i].(string)] = 0
	}

	logger.Infof("[web][statistics] allow site %v", allowSite)
}

func Counter(_ context.Context, c *app.RequestContext) {
	referer := c.GetHeader("Referer")
	if referer == nil {
		c.String(400, "no referer")
		return
	}

	ref := util.BytesToStr(referer)
	if u, err := url.Parse(ref); err != nil {
		c.String(400, "parse referer err")
		return
	} else if _, ok := allowSite[u.Host]; !ok {
		c.String(403, "host %s now allow", u.Host)
		return
	}

	// set to cache

	c.SetContentType("application/javascript")
	c.SetStatusCode(200)
	c.WriteString("console.log('hexo-statistics v0.1.0')")
}

func Get(_ context.Context, c *app.RequestContext) {
	c.JSON(200, model.StatisticsResp{
		SitePv: 0,
		SiteUv: 0,
		PagePv: 0,
	})
}
