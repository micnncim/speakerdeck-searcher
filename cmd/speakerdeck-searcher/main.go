package main

import (
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	sd "github.com/micnncim/speakerdeck-searcher"
)

var (
	username   = kingpin.Arg("username", "Name of user whose stared decks.").Required().String()
	query      = kingpin.Arg("query", "The word for querying decks.").Required().String()
	clearCache = kingpin.Flag("clear-cache", "Clear decks cache.").Bool()
)

func main() {
	kingpin.Parse()

	cache, err := sd.NewCache()
	if err != nil {
		log.Fatal(err)
	}

	cacheExpired := cache.Expired(*username)
	if *clearCache || cacheExpired {
		if err := cache.Delete(*username); err != nil {
			log.Fatal(err)
		}
	}

	var decks []*sd.Deck
	decks, err = cache.Get(*username)
	if err != nil {
		url := sd.StaredURL(*username)
		decks, err = sd.Fetch(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	filtered := sd.FilterByTitle(decks, *query)
	if len(filtered) != 0 {
		fmt.Println(sd.RenderDecks(filtered))
	} else {
		fmt.Println("No result found.")
	}

	if !cacheExpired {
		return
	}

	if err := cache.Put(*username, decks); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched decks have been cached. To clear the cache, use --clear-cache flag.")
}
