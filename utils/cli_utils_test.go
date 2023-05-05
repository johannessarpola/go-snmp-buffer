package utils

import (
	"fmt"
	"testing"
)

func TestPrettyPrintSnmpKey_1(t *testing.T) {
	pf := "snmp_"
	pfa := []byte(pf)
	idx := uint64(99)
	idxa := ConvertToByteArr(idx)

	s := fmt.Sprintf("%s%d", pf, idx)

	k := append(pfa[:], idxa...)

	rs := PrettyPrintSnmpKey([]byte(k), len(pf))

	if rs != s {
		t.Error("Invalid key returned")
	}

}

func TestPrettyPrintSnmpKey_2(t *testing.T) {

	for i := 0; i < 99; i++ {
		pf := RandomString(6)
		pfa := []byte(pf)
		idx := uint64(i)
		idxa := ConvertToByteArr(idx)
		k := append(pfa[:], idxa...)

		s := fmt.Sprintf("%s%d", pf, idx)
		rs := PrettyPrintSnmpKey([]byte(k), len(pf))
		if rs != s {
			t.Error("Invalid key returned")
		}

	}

}
