package utils

import (
	"testing"

	"github.com/adrg/strutil/metrics"
)

func TestRandomString_1(t *testing.T) {

	s1 := RandomString(16)
	s2 := RandomString(16)

	if s1 == s2 {
		t.Error("Same strings generated")
	}
}

func TestRandomString_2(t *testing.T) {
	ham := metrics.NewHamming()
	min_d := 20 // EXtremely unlikely to be closer than 20 modifications from each other with char len of 100
	mx := 100
	for i := 0; i < 64; i++ {
		s1 := RandomString(mx)
		s2 := RandomString(mx)
		d := ham.Distance(s1, s2)
		if d < min_d {
			t.Error("Strings were not random enough")
		}
	}
}
