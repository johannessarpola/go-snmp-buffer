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
- you only need to lock(miten) the next idx generation (cur & offs) 
- have keys in different database and 'service'
- compare for loop and streams perf with badger

# Minor things to do
- Limit the trap output to necessary fields
- [test] Producer needs some variation in the output
- Listener should be constantly running and listening with a cmd
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

9.4.2023 on 2,6 GHz 6-Core Intel Core i7 (mbp)
```sh
#ants:1
Time it took to take in 10000 packets was 555 ms
Time it took to take in 10000 packets was 583 ms
Time it took to take in 10000 packets was 521 ms

# ants: 100
Time it took to take in 10000 packets was 374 ms
Time it took to take in 10000 packets was 376 ms
Time it took to take in 10000 packets was 382 ms

#ants: 1000
Time it took to take in 10000 packets was 378 ms
Time it took to take in 10000 packets was 398 ms
Time it took to take in 10000 packets was 406 ms
```

Space consumed with `level=info msg="Current idx: 808102"` is 84mb (packets are very simple)

Same env with forwarder with same benchmarker

```
Time it took to take in 10000 packets was 287 ms
Time it took to take in 10000 packets was 248 ms
Time it took to take in 10000 packets was 212 ms
```
