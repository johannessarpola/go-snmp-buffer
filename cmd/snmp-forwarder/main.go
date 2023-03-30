package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto/z"
	g "github.com/gosnmp/gosnmp"
	db "github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/snmp"
	"github.com/johannessarpola/go-network-buffer/utils"
)

func main() {
	// TODO Read SNMP from disk -> send forward with some adapter(?)

	data := db.NewData("../../_tmp", "snmp_") // Will cause conflict probably if run with listener
	stream := data.DB.NewStream()

	// -- Optional settings
	stream.NumGo = 8                    // Set number of goroutines to use for iteration.
	stream.Prefix = []byte("snmp_")     // Leave nil for iteration over the whole DB.
	stream.LogPrefix = "snmp.Streaming" // For identifying stream logs. Outputs to Logger.

	// ChooseKey is called concurrently for every key. If left nil, assumes true by default.
	// TODO Filtering does not seem to be working for now
	// stream.ChooseKey = func(item *badger.Item) bool {
	// 	return bytes.HasSuffix(item.Key(), []byte("snmp_"))
	// }

	// Send is called serially, while Stream.Orchestrate is running.
	stream.Send = func(buf *z.Buffer) error {
		list, err := badger.BufferToKVList(buf)
		if err != nil {
			return err
		}
		for _, kv := range list.Kv {
			if kv.StreamDone == true {
				return nil
			}

			// TODO Clean up at some point
			prefix := []byte("snmp_")                        // TODO Ugly as f
			k := utils.ConvertToUint64(kv.Key[len(prefix):]) // TODO Ugly as f
			fmt.Printf("key: %d\n", k)
			v := kv.Value

			p, err := snmp.Decode(v)
			if err != nil {
				log.Println("could not decode!!")
			}
			println(p.Community)
			// TODO Clean up at some point, just for debug for now
			for i, variable := range p.Variables {
				fmt.Printf("%d: oid: %s ", i, variable.Name)

				switch variable.Type {
				case g.OctetString:
					bytes := variable.Value.([]byte)
					fmt.Printf("string: %s\n", string(bytes))
				case g.TimeTicks:
					n := g.ToBigInt(variable.Value)
					tm := time.Unix(n.Int64(), 0)
					fmt.Printf("time: %s\n", tm.String())
				default:
					// ... or often you're just interested in numeric values.
					// ToBigInt() will return the Value as a BigInt, for plugging
					// into your calculations.
					fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
				}
			}

		}
		return nil // TODO Needs some handling
	}

	// Run the stream
	if err := stream.Orchestrate(context.Background()); err != nil {
		log.Fatal("Apuva")
	}

}
