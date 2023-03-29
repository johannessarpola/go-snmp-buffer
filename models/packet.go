package models

import g "github.com/gosnmp/gosnmp"

type Variable interface {
	Value() interface{}
	Name() string
	Type() g.Asn1BER
}

type PacketSubset interface {
	Community() string
	Variables() []Variable
}

type Packet struct {
	Version            g.SnmpVersion
	MsgFlags           g.SnmpV3MsgFlags
	SecurityModel      g.SnmpV3SecurityModel
	SecurityParameters g.SnmpV3SecurityParameters // interface
	ContextEngineID    string
	ContextName        string
	Community          string
	PDUType            g.PDUType
	MsgID              uint32
	RequestID          uint32
	MsgMaxSize         uint32
	Error              g.SNMPError
	ErrorIndex         uint8
	NonRepeaters       uint8
	MaxRepetitions     uint32
	Variables          []g.SnmpPDU
}

func NewPacket(original *g.SnmpPacket) Packet {
	// TODO Cleanup, most are unnecessary
	return Packet{
		Version:            original.Version,
		MsgFlags:           original.MsgFlags,
		SecurityModel:      original.SecurityModel,
		SecurityParameters: original.SecurityParameters,
		ContextEngineID:    original.ContextEngineID,
		ContextName:        original.ContextName,
		Community:          original.Community,
		PDUType:            original.PDUType,
		MsgID:              original.MsgID,
		RequestID:          original.RequestID,
		MsgMaxSize:         original.MsgMaxSize,
		Error:              original.Error,
		ErrorIndex:         original.ErrorIndex,
		NonRepeaters:       original.NonRepeaters,
		MaxRepetitions:     original.MaxRepetitions,
		Variables:          original.Variables,
	}
}
