package utils

import (
	"fmt"
	"time"
)

/*
Measure rate with channel
*/
func MeasureRate(done <-chan bool) {

	measure_interval := 10000 // TODO Maybe configurable
	i := 0
	start := time.Now().UnixMilli()
	fmt.Printf("Star time was %d", start) // TODO

	for range done {
		i++
		if i%measure_interval == 0 {
			now := time.Now().UnixMilli()
			dur := now - start
			fmt.Printf("Time it took to take in %d packets was %d ms\n", measure_interval, dur) // TODO
			start = now
		}
	}

}
