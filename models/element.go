package models

/*
Simple wrapper with some construtors for element stored in the DB
*/
type Element struct {
	// TODO Probably worth to add some meta here
	Value []byte
}

func EmptyElement() Element {
	return Element{
		Value: nil,
	}
}

func NewElement(v []byte) Element {
	return Element{
		Value: v,
	}
}
