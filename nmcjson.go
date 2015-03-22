package nmcjson

import (
	"github.com/btcsuite/btcd/btcjson"
)

// Init registers the NMC-specific commands with btcjson.
func Init() {
	btcjson.RegisterCustomCmd("name_new", NameNewFromRaw, NameNewReplyParse, nmcHelpStrings["name_new"])
	btcjson.RegisterCustomCmd("name_update", NameUpdateFromRaw, NameUpdateReplyParse, nmcHelpStrings["name_update"])
	btcjson.RegisterCustomCmd("name_firstupdate", NameFirstUpdateFromRaw, NameFirstUpdateReplyParse, nmcHelpStrings["name_firstupdate"])
	btcjson.RegisterCustomCmd("name_list", NameListFromRaw, NameListReplyParse, nmcHelpStrings["name_list"])
	btcjson.RegisterCustomCmd("name_history", NameHistoryFromRaw, NameHistoryReplyParse, nmcHelpStrings["name_history"])
	btcjson.RegisterCustomCmd("name_show", NameShowFromRaw, NameShowReplyParse, nmcHelpStrings["name_show"])
	btcjson.RegisterCustomCmd("name_scan", NameScanFromRaw, NameScanReplyParse, nmcHelpStrings["name_scan"])
	btcjson.RegisterCustomCmd("name_filter", NameFilterFromRaw, NameFilterReplyParse, nmcHelpStrings["name_filter"])
}
