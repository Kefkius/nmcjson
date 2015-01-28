package nmcjson

import (
	"encoding/json"
)

// NameShowResult models the data from the name_show command.
type NameShowResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Txid      string `json:"txid"`
	Address   string `json:"address"`
	ExpiresIn int64  `json:"expires_in"`
}

// NameListResult models the data from the name_list command.
type NameListResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Address   string `json:"address"`
	ExpiresIn int64  `json:"expires_in"`
}

// NameHistoryResult models the data from the name_history command.
type NameHistoryResult NameShowResult

// NameScanResult models the data from the name_scan command.
type NameScanResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	ExpiresIn int64  `json:"expires_in"`
}

// NameFilterResult models the data from the name_filter command.
type NameFilterResult NameScanResult

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
