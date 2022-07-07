package main

import (
	"github.com/stulzq/hexo-statistics/common/globalwait"
	"github.com/stulzq/hexo-statistics/job"
	"github.com/stulzq/hexo-statistics/job/uv_renew"
	"github.com/stulzq/hexo-statistics/web"
)

func main() {
	go job.RunAt("UV_Renew", "00:00:01", uv_renew.Work)

	web.Start()

	globalwait.Wait()
}
