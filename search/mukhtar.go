package search

import (
	"bufio"

	"github.com/pkg/errors"
)

var rawMukhtar = "files/mukhtar.txt"

type MukhtarAlsahah struct{}

func (*MukhtarAlsahah) Search(w string) (*SearchResult, error) {
	f, err := rawFiles.Open(rawMukhtar)
	if err != nil {
		err = errors.Wrapf(err, "couldn't open source for %s", rawMukhtar)
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

func init() {
	var t Source = &MukhtarAlsahah{}
	Sources = append(Sources, t)
}
