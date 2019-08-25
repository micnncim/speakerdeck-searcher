package main

import (
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	sd "github.com/micnncim/speaker-deck-searcher"
)

var (
	username   = kingpin.Arg("username", "Name of user whose stared decks.").Required().String()
	query      = kingpin.Arg("query", "The word to use for query decks.").Required().String()
	clearCache = kingpin.Flag("clear-cache", "Clear cache of decks.").Bool()
)

func main() {
	kingpin.Parse()

	cache, err := sd.NewCache()
	if err != nil {
		log.Fatal(err)
	}

	if *clearCache {
		if err := cache.Delete(); err != nil {
			log.Fatal(err)
		}
	}

	var decks []*sd.Deck
	decks, err = cache.Get()
	if err != nil {
		url := sd.StaredURL(*username)
		decks, err = sd.Fetch(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	filtered := sd.FilterByTitle(decks, *query)
	fmt.Println(sd.RenderDecks(filtered))

	if err := cache.Put(decks); err != nil {
		log.Fatal(err)
	}
}
