package nmcjson

// nmcHelpStrings contains the help messages for API calls.
var nmcHelpStrings = map[string]string{
	"name_new": `name_new "name"
Pre-order a new name`,
	"name_update": `name_update "name" "value" [toaddress]
Update and possibly transfer a name`,
	"name_firstupdate": `name_firstupdate "name" "rand" [tx] "value" [toaddress]
Perform a first update after a name_new reservation.
Note that the first update will go into a block 12 blocks after the name_new, at the soonest`,
	"name_filter": `name_filter [regexp] [maxage=36000] [from=0] [nb=0] [stat]
Scan and filter names
[regexp] : apply [regexp] on names, empty means all names
[maxage] : look in last [maxage] blocks
[from] : show results from number [from]
[nb] : show [nb] results, 0 means all
[stat] : show some stats instead of results`,
	"name_history": `name_history "identifier"
    List all name values of a name.`,
	"name_list": `name_list [name]
    List my own names`,
	"name_scan": `name_scan [start-identifier] [max-return=500]
    Scan all identifiers, starting at start-identifier and returning a maximum number of entries`,
	"name_show": `name_show "identifier"
    Show values of a name`,
}
