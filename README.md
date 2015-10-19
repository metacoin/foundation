# foundation

This library provides a simple interface to communicate with a local florincoin client in Go.

### example

An example can be found here: [examples/main.go]

Here's a sample from that code

```
response, err := foundation.RPCCall("getnetworkhashps")
// ... error checking and stuff ...
if networkHashrate, ok := response.(int64); ok {
    fmt.Printf("Florincoin network hashrate: %v\n", networkHashrate)
```

This is the output (remember to export your RPC username and password):

```
$ export F_TOKEN=verystrongrpcpassword
$ export F_USER=florpcuser
$ export F_URI=http://127.0.0.1:18322
$ go run main.go
Florincoin network hashrate: 874124892
$
```
### security

Remember, by exposing an RPC interface you're opening up an attack vector. Disable any wallet functions from foundation or make sure the process that runs it is sufficiently guarded from intruders. Always encrypt your wallet file. This is mostly useful for block explorer type stuff, parsing the block chain and building distributed apps using tx-comment protocols.

### license

[MIT]

[examples/main.go]:examples/main.go
[MIT]:LICENSE.md
