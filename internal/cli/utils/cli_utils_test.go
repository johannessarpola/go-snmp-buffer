package utils

import (
	"fmt"
	"testing"

	c "github.com/johannessarpola/go-network-buffer/pkg/conversions"
	s "github.com/johannessarpola/go-network-buffer/pkg/strings"
)

func TestPrettyPrintSnmpKey_1(t *testing.T) {
	pf := "snmp_"
	pfa := []byte(pf)
	idx := uint64(99)
	idxa := c.ConvertToByteArr(idx)

	s := fmt.Sprintf("%s%d", pf, idx)

	k := append(pfa[:], idxa...)

	rs := PrettyPrintPrefixedKey([]byte(k), len(pf))

	if rs != s {
		t.Error("Invalid key returned")
	}

}

func TestPrettyPrintSnmpKey_2(t *testing.T) {

	for i := 0; i < 99; i++ {
		pf := s.RandomString(6)
		pfa := []byte(pf)
		idx := uint64(i)
		idxa := c.ConvertToByteArr(idx)
		k := append(pfa[:], idxa...)

		s := fmt.Sprintf("%s%d", pf, idx)
		rs := PrettyPrintPrefixedKey([]byte(k), len(pf))
		if rs != s {
			t.Error("Invalid key returned")
		}

	}

}
