package speaker_deck_searcher

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	errors "golang.org/x/xerrors"
)

func Fetch(baseURL string) ([]*Deck, error) {
	var decks []*Deck
	var i int
	for {
		i++
		url := fmt.Sprintf("%s?page=%d", baseURL, i)
		ds, err := fetchSinglePage(url)
		if err != nil {
			return nil, err
		}
		if len(ds) == 0 {
			break
		}
		decks = append(decks, ds...)
	}
	return decks, nil
}

func fetchSinglePage(url string) ([]*Deck, error) {
	doc, err := documentFromURL(url)
	if err != nil {
		return nil, err
	}

	var decks []*Deck
	doc.Find(".row > .col-12").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Find("a").Attr("href")
		if !ok {
			return
		}
		title, ok := s.Find("a").Attr("title")
		if !ok {
			return
		}

		var stars int
		var views int
		s.Find(".py-3").Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				starsStr := strings.TrimSpace(s.Text())
				stars, _ = strconv.Atoi(starsStr)
			}
			s.Find("span").Each(func(_ int, s *goquery.Selection) {
				viewsStr, ok := s.Attr("title")
				if ok {
					viewsStr = strings.Trim(viewsStr, " views")
					viewsStr = strings.Replace(viewsStr, ",", "", -1)
					views, _ = strconv.Atoi(viewsStr)
				}
			})
		})

		decks = append(decks, &Deck{
			Title: title,
			URL:   fmt.Sprintf("%s%s", SpeakerDeckBaseURL, url),
			Stars: stars,
			Views: views,
		})
	})

	return decks, nil
}

func documentFromURL(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("query: status code is not 200")
	}
	return goquery.NewDocumentFromReader(resp.Body)
}
