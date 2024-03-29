package statistics

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/stulzq/hexo-statistics/cache"
	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/keys"
	"github.com/stulzq/hexo-statistics/logger"
	"github.com/stulzq/hexo-statistics/web/handlers/statistics/model"
	"net/url"
)

var allowSite map[string]int

const (
	ERR_PARSE_URL     = "parse referer err"
	ERR_NO_PARAMETER  = "no parameter u"
	ERR_URL_NOT_ALLOW = "host %s not allow"
)

func init() {
	allowSite = config.GetStatAllowSite()
	logger.Infof("[web][statistics] allow site %v", allowSite)
}

func Counter(ctx context.Context, c *app.RequestContext) {
	u, err := checkParam(c)
	if err != nil {
		c.JSON(400, utils.H{"msg": err.Error()})
		return
	}

	siteUvKey := keys.GenSiteUv(u.Host)
	sitePvKey := keys.GenSitePv(u.Host)
	pagePvKey := keys.GenPagePv(u.Host, u.Path)
	// set to cache
	cli := cache.GetClient().Pipeline()
	cli.PFAdd(ctx, siteUvKey, c.ClientIP())
	cli.Incr(ctx, sitePvKey)
	cli.Incr(ctx, pagePvKey)

	if _, err := cli.Exec(ctx); err != nil {
		c.JSON(500, utils.H{"msg": "set to cache err"})
		logger.Error("[Web][Stat][Get] set to cache error", err)
		return
	}

	c.SetContentType("application/javascript")
	c.SetStatusCode(200)
	c.WriteString("console.log('hexo-statistics v0.1.0')")
}

func Get(ctx context.Context, c *app.RequestContext) {
	u, err := checkParam(c)
	if err != nil {
		c.JSON(400, utils.H{"msg": err.Error()})
		return
	}

	siteUvKey := keys.GenSiteUv(u.Host)
	sitePvKey := keys.GenSitePv(u.Host)
	pagePvKey := keys.GenPagePv(u.Host, u.Path)
	siteUvArchKey := keys.GenSiteUvArchive(u.Host)
	// get from cache
	cli := cache.GetClient().Pipeline()
	siteUv := cli.PFCount(ctx, siteUvKey)
	siteArchUv := cli.Get(ctx, siteUvArchKey)
	sitePv := cli.Get(ctx, sitePvKey)
	pagePv := cli.Get(ctx, pagePvKey)
	if _, err := cli.Exec(ctx); err != nil && err != redis.Nil {
		c.JSON(500, utils.H{"msg": "read data err"})
		logger.Error("[Web][Stat][Get] read data error", err)
		return
	}

	var siteUvArchInt, sitePvInt, pagePvInt int64
	if pagePv.Err() == nil {
		pagePvInt, _ = pagePv.Int64()

	}

	if sitePv.Err() == nil {
		sitePvInt, _ = sitePv.Int64()
	}

	if siteArchUv.Err() == nil {
		siteUvArchInt, _ = siteArchUv.Int64()
	}

	c.JSON(200, model.StatisticsResp{
		SitePv: sitePvInt,
		SiteUv: siteUv.Val() + siteUvArchInt,
		PagePv: pagePvInt,
	})
}

func checkParam(c *app.RequestContext) (*url.URL, error) {
	param := c.Query("u")
	if param == "" {
		return nil, errors.New(ERR_NO_PARAMETER)
	}

	var u *url.URL
	var err error
	if u, err = url.Parse(param); err != nil {
		return nil, errors.Wrapf(err, ERR_PARSE_URL)
	} else if _, ok := allowSite[u.Host]; !ok {
		return nil, errors.Errorf(ERR_URL_NOT_ALLOW, u.Host)
	}

	return u, nil
}
