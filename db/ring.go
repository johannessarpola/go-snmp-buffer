package db

type RingBuffer interface {
	/*
		Set the maximum size of the ring buffer.
	*/
	SetCapacity(size uint64)
	/*
		Capacity returns the current capacity of the ring buffer.
	*/
	Capacity() uint64
	/*
		ContentSize returns the current number of elements inside the ring buffer.
	*/
	ContentSize() uint64
	/*
		Enqueue a value into the Ring buffer.
	*/
	Enqueue(i interface{})
	/*
		Dequeue a value from the Ring buffer.

		Returns nil if the ring buffer is empty.
	*/
	Dequeue() interface{}
	/*
		Values returns a slice of all the values in the circular buffer without modifying them at all.
		The returned slice can be modified independently of the circular buffer. However, the values inside the slice
		are shared between the slice and circular buffer.
	*/
	Values() []interface{}
	/*
		Read the value that Dequeue would have dequeued without actually dequeuing it.

		Returns nil if the ring buffer is empty.
	*/
	Peek() interface{}
}
