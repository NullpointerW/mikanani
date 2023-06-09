package test

import (
	"fmt"
	"strings"
	"testing"

	I "github.com/NullpointerW/anicat/crawl/information"
	util "github.com/NullpointerW/anicat/utils"
)

func TestInfoSearch(t *testing.T) {
	I.InfoPageScrape("凉宫春日的消失")
}

func TestInfoScraping(t *testing.T) {
	tip, err := I.Scrape("铃芽之旅")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for k, v := range tip {
		fmt.Println(k)
		fmt.Println(v)
	}
}
func TestBgmTVInfoScrape(t *testing.T) {
	tip, err := I.BgmTVInfoScrape(274234)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for k, v := range tip {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func TestTMDB(t *testing.T) {
	_, d, e := I.FloderSearch(I.TMDB_TYP_TV, "凉宫春日的忧郁")
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
	pd,_:=util.ParseShort02Time(strings.ReplaceAll(d, " ", ""))
	fmt.Println(pd)
}
