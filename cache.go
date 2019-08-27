package speaker_deck_searcher

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultCacheTTL = 14 * 24 * time.Hour
)

type cache struct {
	dir string
	ttl time.Duration
}

type CacheOption func(*cache)

func WithCacheDir(dir string) CacheOption {
	return func(c *cache) {
		c.dir = dir
	}
}

func WithTTL(ttl time.Duration) CacheOption {
	return func(c *cache) {
		c.ttl = ttl
	}
}

func NewCache(opts ...CacheOption) (*cache, error) {
	defaultCacheDir, err := defaultCacheDir()
	if err != nil {
		return nil, err
	}
	c := &cache{
		dir: defaultCacheDir,
		ttl: defaultCacheTTL,
	}
	for _, opt := range opts {
		opt(c)
	}
	if err := os.MkdirAll(c.dir, 0777); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *cache) Get(key string) ([]*Deck, error) {
	f, err := os.Open(filepath.Join(c.dir, key+".json"))
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

func (c *cache) Put(key string, decks []*Deck) error {
	f, err := os.Create(filepath.Join(c.dir, key+".json"))
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(decks)
}

func (c *cache) Delete(key string) error {
	return os.RemoveAll(filepath.Join(c.dir, key+".json"))
}

func (c *cache) Expired(key string) bool {
	info, err := os.Stat(filepath.Join(c.dir, key+".json"))
	if err != nil {
		return true
	}
	return info.ModTime().Add(c.ttl).Before(time.Now())
}

func defaultCacheDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	// TODO: Set cache dir properly.
	return filepath.Join(homeDir, ".cache", "speaker-deck-searcher"), nil
}
