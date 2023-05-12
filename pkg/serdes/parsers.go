package serdes

import (
	"bytes"
	"encoding/gob"

	"github.com/johannessarpola/go-network-buffer/pkg/models"
)

func DecodeGob(arr []byte) (models.Packet, error) {
	var p models.Packet
	buf := bytes.NewBuffer(arr)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&p)
	return p, err
}

func EncodeGob(packet *models.Packet) ([]byte, error) {
	var buf bytes.Buffer // Stand-in
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(packet)
	return buf.Bytes(), err
}
