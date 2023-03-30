# Things to do

- Connect the output forwarder to Kafka and/or Pulsar (through adapter)
    - Add support for Schema Registry with Protobuf so that
    output stream adheres to a schema and reduces the amount of data transfered
- Limit the trap output to necessary fields
- Producer needs some variation in the output
- Listener should be constantly running and listening with a cmd
- Forwarder can be started every with set times like every 10 seconds
    - Forwarder should know on what index it is or either remove sent records
- Persistence should expire records after a set while (e.g. 7 days)


# Minor things to do

- Replace prints with different level (https://www.honeybadger.io/blog/golang-logging/)