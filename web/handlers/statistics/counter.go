package statistics

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/stulzq/hexo-statistics/cache"
	"github.com/stulzq/hexo-statistics/config"
	"github.com/stulzq/hexo-statistics/logger"
	"github.com/stulzq/hexo-statistics/util"
	"github.com/stulzq/hexo-statistics/web/handlers/statistics/model"
	"net/url"
)

var allowSite map[string]int

const (
	ERR_PARSE_REFERER     = "parse referer err"
	ERR_NO_REFERER        = "no referer"
	ERR_REFERER_NOT_ALLOW = "host %s not allow"
)

func init() {
	data := config.Get("statistics:site").([]interface{})
	allowSite = make(map[string]int, len(data))

	for i := 0; i < len(data); i++ {
		allowSite[data[i].(string)] = 0
	}

	logger.Infof("[web][statistics] allow site %v", allowSite)
}

func Counter(ctx context.Context, c *app.RequestContext) {
	u, err := checkReferer(c)
	if err != nil {
		c.JSON(400, utils.H{"msg": err.Error()})
		return
	}

	siteUvKey := genSiteUv(u.Host)
	sitePvKey := genSitePv(u.Host)
	pagePvKey := genPagePv(u.Host, u.Path)
	// set to cache
	fmt.Println(c.ClientIP())
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
	u, err := checkReferer(c)
	if err != nil {
		c.JSON(400, utils.H{"msg": err.Error()})
		return
	}

	siteUvKey := genSiteUv(u.Host)
	sitePvKey := genSitePv(u.Host)
	pagePvKey := genPagePv(u.Host, u.Path)
	// get from cache
	cli := cache.GetClient().Pipeline()
	siteUv := cli.PFCount(ctx, siteUvKey)
	sitePv := cli.Get(ctx, sitePvKey)
	pagePv := cli.Get(ctx, pagePvKey)
	if _, err := cli.Exec(ctx); err != nil && err != redis.Nil {
		c.JSON(500, utils.H{"msg": "read data err"})
		logger.Error("[Web][Stat][Get] read data error", err)
		return
	}

	var sitePvInt, pagePvInt int64
	if pagePv.Err() == nil {
		pagePvInt, _ = pagePv.Int64()

	}

	if sitePv.Err() == nil {
		sitePvInt, _ = sitePv.Int64()
	}

	c.JSON(200, model.StatisticsResp{
		SitePv: sitePvInt,
		SiteUv: siteUv.Val(),
		PagePv: pagePvInt,
	})
}

func genSiteUv(host string) string {
	return fmt.Sprintf("siteuv:%s", host)
}

func genSitePv(host string) string {
	return fmt.Sprintf("sitepv:%s", host)
}

func genPagePv(host string, path string) string {
	return fmt.Sprintf("pagepv:%s:%s", host, path)
}

func checkReferer(c *app.RequestContext) (*url.URL, error) {
	referer := c.GetHeader("Referer")
	if referer == nil {
		return nil, errors.New(ERR_NO_REFERER)
	}

	ref := util.BytesToStr(referer)
	var u *url.URL
	var err error
	if u, err = url.Parse(ref); err != nil {
		return nil, errors.Wrapf(err, ERR_PARSE_REFERER)
	} else if _, ok := allowSite[u.Host]; !ok {
		return nil, errors.Errorf(ERR_REFERER_NOT_ALLOW, u.Host)
	}

	return u, nil
}
