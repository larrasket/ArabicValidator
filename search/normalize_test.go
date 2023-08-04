package search

import "testing"

func TestNormalize(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"Hamza should be normalized", "دإب", "دأب"},
		{"Hamza should be normalized", "دؤب", "دأب"},
		{"Hamza should be normalized", "دؤب", "دأب"},

		{"Kashia and Tashkeel should be normalized", "لَعــب", "لعب"},
		// TODO add more test cases.
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Normalize(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
