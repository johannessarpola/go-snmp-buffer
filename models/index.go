package models

import (
	"github.com/johannessarpola/go-network-buffer/utils"
)

type Index struct {
	Name   string
	Value  uint64
	Exists bool
}

func NewIndex(name string, value uint64) Index {
	return Index{
		Name:   name,
		Value:  value,
		Exists: true,
	}
}

func ZeroIndex(name string) Index {
	return Index{
		Name:   name,
		Value:  0,
		Exists: false,
	}
}

func (idx *Index) AsBytes() ([]byte, []byte) {
	return idx.KeyAsBytes(), idx.ValueAsBytes()
}

func (idx *Index) KeyAsBytes() []byte {
	return []byte(idx.Name)
}

func (idx *Index) ValueAsBytes() []byte {
	return utils.ConvertToByteArr(idx.Value)
}

func (idx *Index) WithValue(value uint64) *Index {
	idx.WithExists(true).Value = value
	return idx
}

func (idx *Index) WithExists(exists bool) *Index {
	idx.Exists = exists
	return idx
}

func (idx *Index) SetValue(value uint64) *Index {
	return idx.WithExists(true).WithValue(value)
}

func (idx *Index) Increment() *Index {
	return idx.SetValue(idx.Value + 1)
}
