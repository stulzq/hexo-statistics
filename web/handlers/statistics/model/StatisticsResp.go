package model

type StatisticsResp struct {
	SitePv int64 `json:"sitePv"`
	SiteUv int64 `json:"siteUv"`
	PagePv int64 `json:"pagePv"`
}
