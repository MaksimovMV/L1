package main

import "fmt"

type message struct {
}

func (m *message) sendMessage(w webService) {
	w.post()
}

type webService interface {
	post()
}

type twitter struct {
}

func (t *twitter) twit() {
	fmt.Println("Twit was published")
}

type facebook struct {
}

func (f *facebook) post() {
	fmt.Println("Facebook post was published")
}

type twitterAdapter struct {
	twitter *twitter
}

func (tA *twitterAdapter) post() {
	tA.twitter.twit()
}

func main() {
	m := &message{}
	f := &facebook{}
	m.sendMessage(f)

	t := &twitter{}
	tA := &twitterAdapter{t}
	m.sendMessage(tA)
}
