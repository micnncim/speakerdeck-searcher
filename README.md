# speakerdeck-searcher

The CLI searcher for [Speaker Deck](https://speakerdeck.com).

Currently supports only searching user stared decks.
This tool searches decks by a query word for titles (e.g. kubernetes) and supports insenstive case.

## Installation

```
$ go get github.com/micnncim/speakerdeck-searcher/cmd/speakerdeck-searcher
```

## Usage

```
$ speakerdeck-searcher --help
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
$ speakerdeck-searcher micnncim kubernetes
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
|                      TITLE                      |                                                 URL                                                 |
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
| Kubernetes Handson                              | https://speakerdeck.com/stormcat24/kubernetes-handson                                               |
| Sustainable Kubernetes                          | https://speakerdeck.com/mumoshu/sustainable-kubernetes                                              |
| Kubernetes manifests                            | https://speakerdeck.com/b4b4r07/kubernetes-manifests-management-and-operation-in-mercari            |
| management                                      |                                                                                                     |
+-------------------------------------------------+-----------------------------------------------------------------------------------------------------+
```
