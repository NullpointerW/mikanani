package crawl

import (
	"net/url"
	"strings"

	CFG "github.com/NullpointerW/anicat/conf"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func NewCollector() *colly.Collector {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{Parallelism: 1})
	SetProxy(c)
	return c
}

func SetProxy(c *colly.Collector) {
	if len(CFG.Env.Proxies) != 0 {
		if p, err := proxy.RoundRobinProxySwitcher(
			CFG.Env.Proxies...,
		); err == nil {
			c.SetProxyFunc(p)
		}
	}
}

func ConstructSearch(s string) (utoa string) {
	a := url.QueryEscape(strings.ReplaceAll(s, " ", "+"))
	utoa = strings.ReplaceAll(a, "%2B", "+")
	return
}
