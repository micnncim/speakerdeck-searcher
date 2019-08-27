package speaker_deck_searcher

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

const spinnerSymbol = 14

type Spinner struct {
	*spinner.Spinner
	text string
}

func NewSpinner(text string) *Spinner {
	return &Spinner{
		Spinner: spinner.New(spinner.CharSets[spinnerSymbol], 100*time.Millisecond),
		text:    text,
	}
}

func (s *Spinner) Start() {
	s.Spinner.Writer = os.Stderr
	s.Spinner.Prefix = "\r"
	if len(s.text) > 0 {
		s.Suffix = " " + s.text
	}
	s.Spinner.Start()
}

func (s *Spinner) Stop() {
	s.Spinner.Stop()
}
