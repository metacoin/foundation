package main

import (
	"fmt"

	"github.com/metacoin/foundation"
)

func main() {

	response, err := foundation.RPCCall("getnetworkhashps")
	if err != nil {
		// this is merely an example. handle your errors please.
		fmt.Println(err.Error())
	}

	// type assertion because flojson returns us an interface{}
	if networkHashrate, ok := response.(int64); ok {
		fmt.Printf("Florincoin network hashrate: %v\n", networkHashrate)
	} else {
		// what did i say about properly handing errors?
		fmt.Println("Type assertion failed.")
	}

}
