package snmpdb

type RingBuffer interface {
	/*
		Set the maximum size of the ring buffer.
	*/
	SetCapacity(size uint64) error
	/*
		Capacity returns the current capacity of the ring buffer.
	*/
	Capacity() (uint64, error)
	/*
		ContentSize returns the current number of elements inside the ring buffer.
	*/
	ContentSize() (uint64, error)
	/*
		Enqueue a value into the Ring buffer.
	*/
	Enqueue(i interface{}) error
	/*
		Dequeue a value from the Ring buffer.

		Returns nil if the ring buffer is empty.
	*/
	Dequeue() (interface{}, error)
	/*
		Values returns a slice of all the values in the circular buffer without modifying them at all.
		The returned slice can be modified independently of the circular buffer. However, the values inside the slice
		are shared between the slice and circular buffer.
	*/
	Values() ([]interface{}, error)
	/*
		Read the value that Dequeue would have dequeued without actually dequeuing it.

		Returns nil if the ring buffer is empty.
	*/
	Peek() (interface{}, error)
}
