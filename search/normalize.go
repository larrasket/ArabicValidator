package search

import (
	"strings"

	. "github.com/01walid/goarabic"
)

// normalizing arabic words according to Nisrean Thalji, Nik Adilah Hanin,
// Yasmin Yacob and Sohair Al-Hakeem, “Corpus for Test, Compare and Enhance
// Arabic Root Extraction Algorithms” International Journal of Advanced Computer
// Science and Applications(IJACSA), 8(5), 2017.
// http://dx.doi.org/10.14569/IJACSA.2017.080529
// rules:
// 1) Remove kasheeda symbol ("_").
// 2) Remove punctuations.
// 3) Remove diacritics.
// 4) Remove non-letters.
// 5) Replace hamza‟s forms ‫ء‬ , ‫آ‬, ‫إ‬ , ‫ا‬ ,‫ة‬ with‫أ‬ .
// 6) Duplicating any letter that has the (Shaddah symbol. " َّ ")
func Normalize(w string) string {
	// 1) Remove kasheeda symbol ("_").
	w = RemoveTatweel(w)
	// 2) Remove punctuations.
	// 4) Remove non-letters.
	w = RemoveAllNonArabicChars(w)
	// 3) Remove diacritics.
	w = RemoveTashkeel(w)
	// 5) Replace hamza‟s forms ء, آ, ؤ, إ, ئ with أ
	normHamz := string("أ")
	r := strings.NewReplacer(
		string("ئ"), normHamz,
		string("إ"), normHamz,
		string("ؤ"), normHamz,
		string("آ"), normHamz,
		string("ء"), normHamz)
	w = r.Replace(w)

	// 6) Duplicating any letter that has the (Shaddah symbol. " َّ ")
	// this is unnecessary since all inputs are root
	return w
}
