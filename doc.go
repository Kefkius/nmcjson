/*
Package nmcjson extends btcjson with Namecoin-specific JSON-RPC API calls.

Each commands requires a RawCmdParser and ReplyParser, as per btcjson. These functions are implemented as <command>FromRaw() and <command>ReplyParse(), respectively.

Usage

Init() must be called before using any calls in nmcjson, as this function registers commands with btcjson.
*/
package nmcjson
