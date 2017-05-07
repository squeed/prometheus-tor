# prometheus-tor
Small Prometheus exporter for the tor daemon.

## Usage
`prometheus-tor (--tor.control-socket /var/run/tor/control | --tor.control-port localhost:9051)`


## Metrics exposed
* `tor_connection_circuits` - the number of open circuits
* `tor_connection_streams` - the number of open streams
* `tor_connection_orconns` - the number of open ORConns
* `tor_traffic_read_bytes`, `tor_traffic_written_bytes` - cumulative inbound and outbound traffic


## TODO
* Let's cheat and assume the CircuitID is sequential. Instant counter!
