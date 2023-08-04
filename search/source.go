package search

import "embed"

type SearchResult struct {
	Found    bool
	Original string // thr original form the string found on the source book
}

type Source interface {
	Search(w string) (*SearchResult, error)
}

var Sources []Source

//go:embed files/*
var rawFiles embed.FS
