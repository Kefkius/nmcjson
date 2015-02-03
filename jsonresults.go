package nmcjson

import (
	"encoding/json"
)

// NameNewResult models the data from the name_new command.
type NameNewResult struct {
	Txid string
	Rand string
}

// NameUpdateResult models the data from the name_new command.
type NameUpdateResult string

// NameFirstUpdateResult models the data from the name_firstupdate command.
type NameFirstUpdateResult string

// NameShowResult models the data from the name_show command.
type NameShowResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Txid      string `json:"txid"`
	Vout      int    `json:"vout",omitempty`
	Address   string `json:"address"`
	Height    int64  `json:"height",omitempty`
	ExpiresIn int64  `json:"expires_in"`
	Expired   bool   `json:"expired",omitempty`
}

// NameListResult models the data from the name_list command.
type NameListResult struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Txid        string `json:"txid",omitempty`
	Vout        int    `json:"vout",omitempty`
	Address     string `json:"address"`
	Height      int64  `json:"height",omitempty`
	ExpiresIn   int64  `json:"expires_in"`
	Expired     bool   `json:"expired",omitempty`
	Transferred bool   `json:"transferred",omitempty"`
}

// NameHistoryResult models the data from the name_history command.
type NameHistoryResult NameShowResult

// NameScanResult models the data from the name_scan command.
type NameScanResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Txid      string `json:"txid",omitempty`
	Vout      int    `json:"vout",omitempty`
	Address   string `json:"address",omitempty`
	Height    int64  `json:"height",omitempty`
	ExpiresIn int64  `json:"expires_in"`
	Expired   bool   `json:"expired",omitempty`
}

// NameFilterResult models the data from the name_filter command.
type NameFilterResult NameScanResult

// NameNewReplyParse is a ReplyParser.
func NameNewReplyParse(msg json.RawMessage) (interface{}, error) {
	var res NameNewResult
	strs := make([]string, 0, 2)
	err := json.Unmarshal(msg, &strs)
	if err != nil {
		return nil, err
	}
	//err := json.Unmarshal(msg, &res)
	//if err != nil {
	//	return nil, err
	//}
	res = NameNewResult{
		Txid: strs[0],
		Rand: strs[1],
	}
	return res, nil
}

// NameUpdateReplyParse is a ReplyParser.
func NameUpdateReplyParse(msg json.RawMessage) (interface{}, error) {
	var res NameUpdateResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameFirstUpdateReplyParse is a ReplyParser.
func NameFirstUpdateReplyParse(msg json.RawMessage) (interface{}, error) {
	var res NameFirstUpdateResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameShowReplyParse is a ReplyParser.
func NameShowReplyParse(msg json.RawMessage) (interface{}, error) {
	var res NameShowResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameListReplyParse is a ReplyParser.
func NameListReplyParse(msg json.RawMessage) (interface{}, error) {
	var res []NameListResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameHistoryReplyParse is a ReplyParser.
func NameHistoryReplyParse(msg json.RawMessage) (interface{}, error) {
	var res []NameHistoryResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameScanReplyParse is a ReplyParser.
func NameScanReplyParse(msg json.RawMessage) (interface{}, error) {
	var res []NameScanResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NameFilterReplyParse is a ReplyParser.
func NameFilterReplyParse(msg json.RawMessage) (interface{}, error) {
	var res []NameFilterResult
	err := json.Unmarshal(msg, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
