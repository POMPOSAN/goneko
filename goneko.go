package goneko

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://www.nyaa.si"

// parse a single tr element to a Result pointer
func parseTr(r *goquery.Selection) *Result {

	res := &Result{}
	res.Details = &details{}

	tds := r.Find("td")
	tds.Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			res.Cat, _ = s.Find("a").Attr("title")
		case 1:
			res.Title = s.Find("a").Last().Text()
			res.ViewUrl = s.Find("a").Last().AttrOr("href", "")
			res.CommentCount, _ = strconv.Atoi(strings.TrimSpace(s.Find("a.comment, a.comments").Text()))
		case 2:
			res.Magnet, _ = s.Find("a").Last().Attr("href")
		case 3:
			res.Size = s.Text()
		case 4:
			ts, _ := s.Attr("data-timestamp")
			res.Timestamp, _ = strconv.Atoi(ts)
		case 5:
			res.Seeders, _ = strconv.Atoi(s.Text())
		case 6:
			res.Leechers, _ = strconv.Atoi(s.Text())
		case 7:
			res.Completed, _ = strconv.Atoi(s.Text())
		}
	})
	return res
}

// return an array (75) of Result pointers or nil
func Parse(link string) ([]*Result, error) {

	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make([]*Result, 75)
	wg := sync.WaitGroup{}

	trs := doc.Find(".default, .success, .danger")

	trs.Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func(i int, s *goquery.Selection) {
			res := parseTr(s)
			results[i] = res
			wg.Done()
		}(i, s)
	})

	wg.Wait()
	return results, nil
}

// parse the nyaa Home Page
func HomePage() ([]*Result, error) {
	return Parse(baseURL)
}

func Search(opts *Opts) ([]*Result, error) {
	var surl string
	if opts.User == "" {
		surl = fmt.Sprintf("https://nyaa.si/?f=%d&c=%d_%d&q=%s&page=%d", opts.Filter, opts.Cat, opts.Subcat, url.QueryEscape(opts.Query), opts.Page)
	} else {
		surl = fmt.Sprintf("https://nyaa.si/user/%s?f=%d&c=%d_%d&q=%s&page=%d", opts.User, opts.Filter, opts.Cat, opts.Subcat, url.QueryEscape(opts.Query), opts.Page)
	}

	return Parse(surl)
}
