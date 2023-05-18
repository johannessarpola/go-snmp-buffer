package serdes

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

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

func DecodeJson(arr []byte) (models.Packet, error) {
	var p models.Packet
	err := json.Unmarshal(arr, &p)
	return p, err
}

func EncodeJson(packet *models.Packet) ([]byte, error) {
	b, err := json.Marshal(packet)
	return b, err
}
