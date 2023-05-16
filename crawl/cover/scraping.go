package cover

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	CR "github.com/NullpointerW/mikanani/crawl"
	"github.com/NullpointerW/mikanani/errs"
	"github.com/NullpointerW/mikanani/util"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
)

func TouchCoverImg(fpath, cover string) (err error) {
	u, err := coverImgScrape(cover)
	if err != nil {
		return err
	}
	c := CR.NewCollector()
	c.SetRequestTimeout(5 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"

		r.Headers.Set("User-Agent", agent)
		r.Headers.Set("Sec-Ch-Ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		r.Headers.Set("Sec-Ch-Ua-Platform", `"Android"`)
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?1")

		util.Debugf("%#+v", r.Headers)
	})
	c.OnResponse(func(r *colly.Response) {
		exp := coverXpathExp
		doc, e := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if e != nil {
			err = e
			return
		}
		a := htmlquery.FindOne(doc, exp)
		m := htmlquery.InnerText(a)
		dl := strings.ReplaceAll(m, `/m/`, `/l/`)
		fmt.Println("cover file url:", dl)
		//download

		resp, e := http.Get(dl)
		if e != nil {
			err = e
			return
		}

		f, e := os.Create(fpath)
		if e != nil {
			err = e
			return
		}
		defer resp.Body.Close()
		defer f.Close()
		wn, e := io.Copy(f, resp.Body)
		log.Printf("cover file downloaded size:%d", wn)
		if e != nil {
			err = e
			return
		}
		if wn == 0 {
			err = errs.ErrCoverDownLoadZeroSize
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, e error) {
		err = e
	})

	c.Visit(u)

	return err
}

func coverImgScrape(coverName string) (cUrl string, err error) {
	c := CR.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		jsonstr := string(r.Body)
		subjUrl := gjson.Get(jsonstr, "0").Get("url").String()
		u, _ := url.Parse(subjUrl)
		u.RawQuery = ""
		cUrl = u.String() + `photos?type=R`
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, e error) {
		log.Println("Something went wrong:", e)
		err = e
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("coverScrapUrl=%s \n", cUrl)
	})

	parseParam := CR.ConstructSearch(coverName)
	c.Visit(fmt.Sprintf(coverSearchUrl, parseParam))

	return
}
