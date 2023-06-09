# Things to do

- Connect the forwarder to Kafka and/or Pulsar (through adapter)
    - Add support for Schema Registry with Protobuf so that
    output stream adheres to a schema and reduces the amount of data transfered
- Forwarder can be started every with set times like every 10 seconds
    - Forwarder should know on what index it is or either remove sent records
- Persistence should expire records after a set while (e.g. 7 days)
- add a cli to manage the database
- improve benchmarking to be more valid
- Probably could use the inbuilt protobuf (badgerpb) instead of gob
- compare for loop and streams perf with badger

# Minor things to do
- Limit the trap output to necessary fields
- [test] Producer needs some variation in the output
- add cli for producer through args etc
- unify loggers https://stackoverflow.com/questions/29538668/logging-globally-across-packages


# Bugs

(none ofc, software is perfect)

- forwarder does not consume the whole up until latest idx, or fails to save last consumed idx

```log
time="2023-04-11T16:59:53+03:00" level=info msg="Current idx: 65458"
time="2023-04-11T16:59:53+03:00" level=info msg="Offset idx: 65457"
```

steps: take in some traps with listener, start forwarder

# Benchmarks

## V2c

### With encoding/gob


18.5.2023 on CPU AMD Ryzen 7 3700X 8-Core Processor @3.60 GHz

Listener performance with encoding/json (encode)

```sh
# ants: 100
Time it took to take in 10000 packets was 610 ms
Time it took to take in 10000 packets was 625 ms
Time it took to take in 10000 packets was 591 ms
Time it took to take in 10000 packets was 621 ms
Time it took to take in 10000 packets was 609 ms
Time it took to take in 10000 packets was 603 ms

#ants: 1000
Time it took to take in 10000 packets was 586 ms
Time it took to take in 10000 packets was 642 ms
Time it took to take in 10000 packets was 616 ms
Time it took to take in 10000 packets was 652 ms
Time it took to take in 10000 packets was 589 ms
Time it took to take in 10000 packets was 620 ms
Time it took to take in 10000 packets was 583 ms
Time it took to take in 10000 packets was 647 ms
```

Same with forwarder (decode)

```sh
Time it took to take in 10000 packets was 359 ms
Time it took to take in 10000 packets was 353 ms
Time it took to take in 10000 packets was 351 ms
```


Space consumed with `level=info msg="Current idx: 808102"` is 84mb (packets are very simple)




### With encoding/json

18.5.2023 on CPU AMD Ryzen 7 3700X 8-Core Processor @3.60 GHz

Listener performance with encoding/json (encode)

```sh
#ants: 100
Time it took to take in 10000 packets was 659 ms
Time it took to take in 10000 packets was 626 ms
Time it took to take in 10000 packets was 606 ms
Time it took to take in 10000 packets was 625 ms
Time it took to take in 10000 packets was 629 ms
Time it took to take in 10000 packets was 635 ms
Time it took to take in 10000 packets was 634 ms
Time it took to take in 10000 packets was 647 ms

#ants: 1000
Time it took to take in 10000 packets was 612 ms
Time it took to take in 10000 packets was 603 ms
Time it took to take in 10000 packets was 631 ms
Time it took to take in 10000 packets was 615 ms
Time it took to take in 10000 packets was 611 ms
Time it took to take in 10000 packets was 635 ms

```

Same with forwarder (decode)
```sh

Time it took to take in 10000 packets was 328 ms
Time it took to take in 10000 packets was 355 ms
Time it took to take in 10000 packets was 318 ms
Time it took to take in 10000 packets was 333 ms

```

### Conclusion

With the hardware it seems that there is only very marginal performance benefit of using GOB even though it probably
saves space on disk, but the benefit of JSON's portability is greater. 