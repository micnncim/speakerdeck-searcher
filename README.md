# speaker-deck-searcher

The CLI searcher for [Speaker Deck](https://speakerdeck.com).

Currently support only searching user stared decks.
This tool searches decks by a query word for titles (e.g. Kubernetes).

## Installation

```
$ go get github.com/micnncim/speaker-deck-searcher/cmd/speaker-deck-searcher
```

## Usage

```
$ speaker-deck-searcher --help
usage: main [<flags>] <username> <query>

Flags:
  --help         Show context-sensitive help (also try --help-long and --help-man).
  --clear-cache  Clear cache of decks.

Args:
  <username>  Name of user whose stared decks.
  <query>     The word to use for query decks.
```

## Example

```
$ speaker-deck-searcher micnncim Kubernetes
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
|                      TITLE                      |                                                 URL                                                 |
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
| Kubernetes Handson                              | https://speakerdeck.com/stormcat24/kubernetes-handson                                               |
| Sustainable Kubernetes                          | https://speakerdeck.com/mumoshu/sustainable-kubernetes                                              |
| Kubernetes manifests                            | https://speakerdeck.com/b4b4r07/kubernetes-manifests-management-and-operation-in-mercari            |
| management                                      |                                                                                                     |
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
```
