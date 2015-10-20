package foundation

import (
	"errors"
	"os"

	"github.com/metacoin/flojson"
)

func RPCCall(command string, args ...interface{}) (interface{}, error) {
	// create a new command based on the command string given
	// handle errors, and return results to main function
	id := 1
	var cmd flojson.Cmd
	var err error
	var ret interface{}

	switch command {

	case "getblock":
		cmd, err = flojson.NewGetBlockCmd(id, args[0].(string))
	case "getblockhash":
		cmd, err = flojson.NewGetBlockHashCmd(id, args[0].(int))
	case "getblockcount":
		cmd, err = flojson.NewGetBlockCountCmd(id)
	case "getrawtransaction":
		cmd, err = flojson.NewGetRawTransactionCmd(id, args[0].(string))
	case "decoderawtransaction":
		cmd, err = flojson.NewDecodeRawTransactionCmd(id, args[0].(string))
	case "getnewaddress":
		cmd, err = flojson.NewGetNewAddressCmd(id)
	case "getmininginfo":
		cmd, err = flojson.NewGetMiningInfoCmd(id)
	case "getconnectioncount":
		cmd, err = flojson.NewGetConnectionCountCmd(id)
	case "getnetworkhashps":
		cmd, err = flojson.NewGetNetworkHashPSCmd(id)
	case "settxfee":
		cmd, err = flojson.NewSetTxFeeCmd(id, args[0].(int64))
	case "sendtoaddress":
		cmd, err = flojson.NewSendToAddressCmd(id, args[0].(string), args[1].(int64), args[2], args[3], args[4])
	case "walletpassphrase":
		cmd, err = flojson.NewWalletPassphraseCmd(id, args[0].(string), args[1].(int))
	case "signmessage":
		cmd, err = flojson.NewSignMessageCmd(id, args[0].(string), args[1].(string))
	case "verifymessage":
		cmd, err = flojson.NewVerifyMessageCmd(id, args[0].(string), args[1].(string), args[2].(string))
	case "validateaddress":
		cmd, err = flojson.NewValidateAddressCmd(id, args[0].(string))
	case "listtransactions":
		cmd, err = flojson.NewListTransactionsCmd(id)
	case "getbalance":
		cmd, err = flojson.NewGetBalanceCmd(id)
	default:
		err = errors.New("command " + command + " not found")
	}

	if err != nil {
		return ret, err
	}

	if cmd != nil {
		reply, err := SendCommand(cmd)
		if err != nil {
			return ret, err
		} else {
			return reply, nil
		}
	}
	return ret, errors.New("flojson replied with empty NewCmd struct")
}

func SendCommand(cmd flojson.Cmd) (interface{}, error) {

	var ret interface{}
	ftoken := os.Getenv("F_TOKEN")
	if ftoken == "" {
		return ret, errors.New("F_TOKEN environment variable not set")
	}

	fuser := os.Getenv("F_USER")
	if ftoken == "" {
		return ret, errors.New("F_USER environment variable not set")
	}

	furi := os.Getenv("F_URI")
	if furi == "" || len(furi) < 1 {
		//return ret, errors.New("F_URI environment variable not set")
		furi = "127.0.0.1:18322"
	}

	// send command to RPC, get a response in reply
	reply, err := flojson.RpcSend(fuser, ftoken, furi, cmd)
	if err != nil {
		return ret, err
	}

	if reply.Result != nil {
		return reply.Result, nil
	} else {
		// if for some reason the result is nil, but err isn't nil, return a generic warning message
		return ret, errors.New("didn't get a response from the RPC - check auth and permission settings")
	}
}
