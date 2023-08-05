package search

import (
	"bufio"
	"embed"
	"path/filepath"

	"github.com/pkg/errors"
)

type SearchResult struct {
	Found    bool
	Original string // thr original form the string found on the source book
}

type Source struct {
	Search    func(w string) (*SearchResult, error)
	Name, Raw string
	Year      uint16
}

var Sources []*Source

//go:embed files/*
var rawFiles embed.FS

// TODO write test
func TextSearch(r, w string) (*SearchResult, error) {
	f, err := rawFiles.Open(filepath.Join("files", r))
	if err != nil {
		err = errors.Wrapf(err, "couldn't open source for %s", r)
		return nil, err
	}
	defer f.Close()

	s := SearchResult{}

	w = Normalize(w)
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		if Normalize(buf.Text()) == w {
			s.Found = true
			s.Original = buf.Text()
			return &s, nil
		}
	}
	return &s, nil
}

func NewNromalizedSource(name string, year uint16, raw string) *Source {
	s := Source{
		Name: name,
		Year: year,
		Raw:  raw,
	}

	s.Search = func(w string) (*SearchResult, error) {
		return TextSearch(s.Raw, w)
	}
	return &s
}

func init() {
	Sources = []*Source{
		NewNromalizedSource("Mukhtar Alsahah", 1267, "mukhtar.txt"),
		NewNromalizedSource("al-Misbah al-munir fi gharib al-sharh al-kabir", 1368, "misbah.txt"),
		NewNromalizedSource("Taj al-ʿArus Min Jawahir al-Qamus", 1790, "taj.txt"),
		NewNromalizedSource("Asas al-Balagha", 1143, "asas.txt"),
	}
}
