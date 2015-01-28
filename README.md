nmcjson
=======

nmcjson extends the [btcjson](https://github.com/btcsuite/btcjson) package with Namecoin-specific JSON-RPC commands.

## Examples

```Go
        // Register nmcjson commands with btcjson
        nmcjson.Init()
        id := 1
        // Create the command
        nameListCmd, err := nmcjson.NewNameListCmd(id)
        if err != nil {
            panic("Something wrong")
        }
        // Send the command and get a reply
        nameListReply, err := btcjson.RpcSend(myName, myPass, myURL, nameListCmd)
        if err != nil {
            panic("Something wrong")
        }
        // Print the reply
        if nameListReply.Result != nil {
            if info, ok := nameListReply.Result.([]nmcjson.NameListResult); ok {
                for _, v := range info {
                    fmt.Printf("Name: %v\n", v.Name)
                }
            }
        }
```


## Installation

```bash
$ go get github.com/kefkius/nmcjson
```

## TODO

- Identifier registration commands

## License

nmcjson is licensed under the liberal ISC License, the same license btcjson is licensed under.
