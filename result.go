package goneko

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Cat          string
	Title        string
	CommentCount int
	Magnet       string
	ViewUrl      string
	Size         string
	Timestamp    int
	Seeders      int
	Leechers     int
	Completed    int
	Details      *details
}

type details struct {
	Submitter   string
	Information string
	Description string
	InfoHash    string
	Comments    []*comment
}

type comment struct {
	Submitter string
	Timestamp int
	Content   string
}

// get details of the result (/!\ Make a new request to Nyaa)
func (r *Result) GetDetails() error {
	resp, err := http.Get(baseURL + r.ViewUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	rows := doc.Find(".row")
	rows.Each(func(i int, s *goquery.Selection) {
		switch i {
		case 1:
			r.Details.Submitter = s.Find(".col-md-5").First().Children().First().Text()
		case 2:
			r.Details.Information = s.Find(".col-md-5").First().Children().First().Text()
		case 4:
			r.Details.InfoHash = s.Find(".col-md-5").First().Children().Last().Text()
		}
	})

	r.Details.Description, err = doc.Find("#torrent-description").Html()
	if err != nil {
		return err
	}

	doc.Find("#collapse-comments").Children().Each(func(i int, s *goquery.Selection) {
		c := &comment{}
		c.Submitter = s.Find(".col-md-2").Children().Children().First().Text()
		c.Timestamp, _ = strconv.Atoi(s.Find(".comment-details").Children().First().Children().First().AttrOr("data-timestamp", "0"))
		c.Content, _ = s.Find(".comment-content").Html()
		r.Details.Comments = append(r.Details.Comments, c)
	})

	return nil
}
