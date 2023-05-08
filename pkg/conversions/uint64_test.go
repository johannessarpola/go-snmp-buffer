package serdes

import (
	"math"
	"math/rand"
	"testing"
)

func TestConversions_1(t *testing.T) {

	for i := 0; i < 64; i++ {
		n := rand.Uint64()
		barr := ConvertToByteArr(n)
		nn := ConvertToUint64(barr)

		if n != nn {
			t.Error("Numbers were invalid")
		}
	}

}

func TestConversions_2(t *testing.T) {
	const maxuint = math.MaxUint64

	barr := ConvertToByteArr(maxuint)
	nn := ConvertToUint64(barr)

	if maxuint != nn {
		t.Error("Numbers were invalid")
	}

}
