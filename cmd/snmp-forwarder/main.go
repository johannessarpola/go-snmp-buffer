package main

import (
	"fmt"
	"log"
	"time"

	g "github.com/gosnmp/gosnmp"
	db "github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/serdes"
	u "github.com/johannessarpola/go-network-buffer/utils"
	"github.com/panjf2000/ants/v2"
	//"github.com/sirupsen/logrus"
)

//var logger = logrus.New()

func handle_var(variable g.SnmpPDU) {
	fmt.Printf("oid: %s ", variable.Name)

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

func process_element(in *models.Element, print bool) {
	decoded, err := serdes.Decode(in.Value)
	if err != nil {
		log.Println("could not decode!!")
	}

	if print {
		println(decoded.Community)
	}
	for _, variable := range decoded.Variables {
		if print {
			handle_var(variable)
		}
	}
}

func main() {
	// TODO Read SNMP from disk -> send forward with some adapter(?)

	data := db.NewDatabase("../../_tmp", "snmp_") // Will cause conflict probably if run with listener
	//stream := data.GetOffsettedStream(8, "snmp.Forwarder")

	defer ants.Release()
	pool, err := ants.NewPool(100)
	if err != nil {
		panic(err) // TODO
	}
	dones := make(chan bool)
	defer close(dones)
	go u.MeasureRate(dones)
	s, _ := data.RingDB.ContentSize()
	i := 0
	fmt.Printf("Total at %d\n", s)
	d, _ := data.RingDB.Dequeue()
	for d != nil {
		d, _ = data.RingDB.Dequeue()
		i++
		pool.Submit(func() {
			// TODO Remove
			if i%5000 == 0 {
				fmt.Printf("Currently processed %d elements\n", i)
				cs, _ := data.RingDB.ContentSize()
				fmt.Printf("Offset index at %d\n", cs)
			}
			process_element(d, false)
			dones <- true
		})

	}

	// 			println(p.Community)

	// // TODO Refactor to use RingDB Deque
	// // Send is called serially, while Stream.Orchestrate is running.
	// stream.Send = func(buf *z.Buffer) error {
	// 	list, err := badger.BufferToKVList(buf)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	//var last_k uint64 = 0
	// 	for _, kv := range list.Kv {
	// 		if kv.StreamDone {
	// 			logger.Info("Batch done!")
	// 		}

	// 		// TODO Clean up at some point
	// 		prefix := []byte("snmp_")                    // TODO Ugly as f
	// 		k := u.ConvertToUint64(kv.Key[len(prefix):]) // TODO Ugly as f
	// 		fmt.Printf("key: %d\n", k)
	// 		v := kv.Value
	// 		//last_k = k

	// 		p, err := serdes.Decode(v)
	// 		if err != nil {
	// 			log.Println("could not decode!!")
	// 		}
	// 		println(p.Community)
	// 		// TODO Clean up at some point, just for debug for now
	// 		for i, variable := range p.Variables {
	// 			handle_var(variable)
	// 		}

	// 	}
	// 	// data.UpdateOffset(last_k) // TODO Not working anymore
	// 	return nil // TODO Needs some handling
	// }

	// // Run the stream
	// if err := stream.Orchestrate(context.Background()); err != nil {
	// 	log.Fatal("Apuva")
	// }

}
