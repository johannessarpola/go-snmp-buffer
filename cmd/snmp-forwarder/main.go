package main

import (
	"fmt"
	"log"

	db "github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/models"
	"github.com/johannessarpola/go-network-buffer/serdes"
	"github.com/johannessarpola/go-network-buffer/utils"
	u "github.com/johannessarpola/go-network-buffer/utils"
	"github.com/panjf2000/ants/v2"
	//"github.com/sirupsen/logrus"
)

//var logger = logrus.New()

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
			u.PrintVars(variable)
		}
	}
}

func main() {
	// TODO Read SNMP from disk -> send forward with some adapter(?)

	idx_fs, err := utils.NewFileStore("../../_idxs")
	if err != nil {
		log.Fatal("could not open index filestore")
	}
	snmp_fs, err := utils.NewFileStore("../../_snmp")
	if err != nil {
		log.Fatal("could not open snmp filestore")
	}

	defer idx_fs.Close()
	defer snmp_fs.Close()

	idx_db := db.NewIndexDB(idx_fs)                     // TODO prefix?
	snmp_data := db.NewSnmpDB(snmp_fs, idx_db, "snmp_") // TODO Configurable prefix

	defer ants.Release()
	pool, err := ants.NewPool(100)
	if err != nil {
		panic(err) // TODO
	}
	dones := make(chan bool)
	defer close(dones)
	go u.MeasureRate(dones)
	s, _ := snmp_data.Buffer.ContentSize()
	i := 0
	fmt.Printf("Total at %d\n", s)
	d, _ := snmp_data.Buffer.Dequeue()
	for d != nil {
		d, _ = snmp_data.Buffer.Dequeue()
		i++
		err := pool.Submit(func() {
			// TODO Remove
			if i%5000 == 0 {
				fmt.Printf("Currently processed %d elements\n", i)
				cs, _ := snmp_data.Buffer.ContentSize()
				fmt.Printf("Offset index at %d\n", cs)
			}
			process_element(d, false)
			dones <- true
		})
		if err != nil {
			log.Fatal("Could not submit into pool", err)
		}

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
