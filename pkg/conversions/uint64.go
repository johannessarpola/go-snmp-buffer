package serdes

import "encoding/binary"

func ConvertToUint64(buf []byte) uint64 {
	// TODO Add some handling for invalid format
	return binary.BigEndian.Uint64(buf)
}

func ConvertToByteArr(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}
