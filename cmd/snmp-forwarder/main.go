package main

import (
	"bytes"

	"github.com/dgraph-io/badger"
)

func main() {
	// TODO Read SNMP from disk -> send forward with some adapter(?)
	print("Hello world")
	stream := db.NewStream()
	// db.NewStreamAt(readTs) for managed mode.

	// -- Optional settings
	stream.NumGo = 16                     // Set number of goroutines to use for iteration.
	stream.Prefix = []byte("some-prefix") // Leave nil for iteration over the whole DB.
	stream.LogPrefix = "Badger.Streaming" // For identifying stream logs. Outputs to Logger.

	// ChooseKey is called concurrently for every key. If left nil, assumes true by default.
	stream.ChooseKey = func(item *badger.Item) bool {
		return bytes.HasSuffix(item.Key(), []byte("er"))
	}

}
