package ArabicValidtor

import (
	"errors"

	. "github.com/larrasket/ArabicValidator/search"
)

// returns search results from definied sources and an error value which is a
// joined errors (if any).
func Search(w string) ([]*SearchResult, error) {
	var errs error
	var res []*SearchResult
	for _, v := range Sources {
		r, err := v.Search(w)
		if err != nil {
			errs = errors.Join(err)
			continue
		}
		res = append(res, r)
	}
	return res, errs
}
