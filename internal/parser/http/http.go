package http

import (
	"github.com/PuerkitoBio/goquery"
	"golang-modules/pkg/helpers"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Url          string
	Body         io.ReadCloser
	Len          int
	CollectLinks []string
}

func SendRequest(url string) Request {
	res, err := http.Get(url)
	helpers.Check(err)

	return Request{Url: url, Body: res.Body}
}

func (r Request) GetLinks() Request {
	var links []string
	doc, err := goquery.NewDocumentFromReader(r.Body)
	helpers.Check(err)

	doc.Find("body a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		if strings.Index(link, r.Url) == 0 {
			l, _ := url.Parse(link)
			link = l.Scheme + "://" + l.Hostname() + l.Path

			ok := true
			for _, li := range links {
				if li == link {
					ok = false
					break
				}
			}

			if ok {
				links = append(links, link)
			}
		}
	})

	r.CollectLinks = links

	return r

}
