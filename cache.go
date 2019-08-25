package speaker_deck_searcher

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type cache struct {
	file string
}

func NewCache() (*cache, error) {
	file, err := cacheFile()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(file), 0777); err != nil {
		return nil, err
	}
	return &cache{
		file: file,
	}, nil
}

func (c *cache) Get() ([]*Deck, error) {
	f, err := os.Open(c.file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ret []*Deck
	if err := json.NewDecoder(f).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *cache) Put(decks []*Deck) error {
	f, err := os.Create(c.file)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(decks)
}

func (c *cache) Delete() error {
	return os.RemoveAll(c.file)
}

func cacheFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	// TODO: Set cache file properly.
	return filepath.Join(homeDir, ".config", "speaker-deck-searcher", "cache.json"), nil
}
