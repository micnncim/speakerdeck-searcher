package speaker_deck_searcher

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const SpeakerDeckBaseURL = "https://speakerdeck.com"

type Deck struct {
	Title string
	URL   string
	Stars int
	Views int
}

func StaredURL(username string) string {
	return fmt.Sprintf("%s/%s/%s", SpeakerDeckBaseURL, username, "stars")
}

func FilterByTitle(decks []*Deck, pattern string) []*Deck {
	var ret []*Deck
	for _, d := range decks {
		title := strings.ToUpper(d.Title)
		p := strings.ToUpper(pattern)
		if strings.Contains(title, p) {
			ret = append(ret, d)
		}
	}
	return ret
}

func RenderDecks(decks []*Deck) string {
	rows := make([][]string, 0, len(decks))
	for _, d := range decks {
		var title string
		r := []rune(d.Title)
		if len(r) > 32 {
			title = string(r[:32])
		} else {
			title = d.Title
		}
		rows = append(rows, []string{title, d.URL})
	}

	b := &bytes.Buffer{}
	t := tablewriter.NewWriter(b)
	t.SetHeader([]string{"title", "url"})
	t.AppendBulk(rows)
	t.Render()

	return b.String()
}
