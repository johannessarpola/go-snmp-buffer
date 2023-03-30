package models

type Index struct {
	name   string
	value  uint64
	exists bool
}

func NewIndex(name string, value uint64) Index {
	return Index{
		name:   name,
		value:  value,
		exists: true,
	}
}

func EmptyIndex(name string) Index {
	return Index{
		name:   name,
		value:  0,
		exists: false,
	}
}

func (idx *Index) WithValue(value uint64) *Index {
	idx.value = value
	return idx
}

func (idx *Index) WithExists(exists bool) *Index {
	idx.exists = exists
	return idx
}

func (idx *Index) SetValue(value uint64) *Index {
	return idx
		.WithValue(value)
		.WithExists(true)
}