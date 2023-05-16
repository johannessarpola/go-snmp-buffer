package snmp

import (
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
)

type SnmpProtocol uint8

const (
	V1 SnmpProtocol = iota
	V2
	V3
	invalid
)

type HasSnmpProtocol struct {
	Protocol SnmpProtocol `json:"protocol"`
}

func (p SnmpProtocol) String() string {
	switch p {
	case V1:
		return "V1"
	case V2:
		return "V2"
	case V3:
		return "V3"
	default:
		return "Invalid"
	}
}

func SnmpProtocolFromString(s string) (SnmpProtocol, error) {
	switch s {
	case "V1":
		return V1, nil
	case "V2":
		return V2, nil
	case "V3":
		return V3, nil
	default:
		return invalid, errors.New("invalid snmpProtocol")
	}
}

func (p SnmpProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *SnmpProtocol) UnmarshalJSON(data []byte) error {
	var protoString string
	if err := json.Unmarshal(data, &protoString); err != nil {
		return err
	}
	prot, err := SnmpProtocolFromString(protoString)
	*p = prot
	return err
}

func DetermineProtocol(content string) (SnmpProtocol, error) {
	var hasSnmpProtocol HasSnmpProtocol
	if err := json.Unmarshal([]byte(content), &hasSnmpProtocol); err != nil {
		logrus.Error("Error unmarshaling:", err)
		return invalid, err
	}

	return hasSnmpProtocol.Protocol, nil
}
