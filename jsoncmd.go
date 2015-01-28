package nmcjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcjson"
)

// NameShowCmd is a type handling custom marshaling and
// unmarshaling of name_show JSON RPC commands.
type NameShowCmd struct {
	id   interface{}
	Name string
}

// Enforce that NameShowCmd satisfies the Cmd interface.
var _ btcjson.Cmd = &NameShowCmd{}

// NewNameShowCmd creates a new NameShowCmd.
func NewNameShowCmd(id interface{}, name string) (*NameShowCmd, error) {
	return &NameShowCmd{
		id:   id,
		Name: name,
	}, nil
}

// NameShowFromRaw is a RawCmdParser.
func NameShowFromRaw(rawCmd *btcjson.RawCmd) (btcjson.Cmd, error) {
	var nameStr string
	err := json.Unmarshal(rawCmd.Params[0], &nameStr)
	if err != nil {
		return NameShowCmd{}, err
	}

	newCmd, err := NewNameShowCmd(rawCmd.Id, nameStr)
	if err != nil {
		return *newCmd, err
	}
	return *newCmd, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (cmd NameShowCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd NameShowCmd) Method() string {
	return "name_show"
}

// MarshalJSON returns the JSON encoding of cmd. Part of the Cmd interface.
func (cmd NameShowCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		cmd.Name,
	}

	raw, err := btcjson.NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (cmd NameShowCmd) UnmarshalJSON(b []byte) error {
	var r btcjson.RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return btcjson.ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("name parameter must be a string: %v", err)
	}

	newCmd, err := NewNameShowCmd(r.Id, name)
	if err != nil {
		return err
	}
	cmd = *newCmd
	return nil
}

// NameListCmd is a type handling custom marshaling and
// unmarshaling of name_list JSON RPC commands.
type NameListCmd struct {
	id         interface{}
	Identifier string
}

// Enforce that NameListCmd satisfies the Cmd interface.
var _ btcjson.Cmd = &NameListCmd{}

// NewNameListCmd creates a new NameListCmd.
func NewNameListCmd(id interface{}, optArgs ...interface{}) (*NameListCmd, error) {
	var identifier string
	if len(optArgs) > 1 {
		return nil, btcjson.ErrWrongNumberOfParams
	}
	if len(optArgs) > 0 {
		a, ok := optArgs[0].(string)
		if !ok {
			return nil, errors.New("optional argument identifier is not a string")
		}
		identifier = a
	}
	return &NameListCmd{
		id:         id,
		Identifier: identifier,
	}, nil
}

// NameListFromRaw is a RawCmdParser.
func NameListFromRaw(rawCmd *btcjson.RawCmd) (btcjson.Cmd, error) {
	var nameStr string
	params := make([]interface{}, 0, 1)
	if len(rawCmd.Params) > 0 {
		err := json.Unmarshal(rawCmd.Params[0], &nameStr)
		if err != nil {
			return NameListCmd{}, err
		}
		params = append(params, nameStr)
	}
	return NewNameListCmd(rawCmd.Id, params)
}

// Id satisfies the Cmd interface by returning the id of the command.
func (cmd NameListCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd NameListCmd) Method() string {
	return "name_list"
}

// MarshalJSON returns the JSON encoding of cmd. Part of the Cmd interface.
func (cmd NameListCmd) MarshalJSON() ([]byte, error) {
	params := make([]interface{}, 0, 1)
	if cmd.Identifier != "" {
		params = append(params, cmd.Identifier)
	}

	raw, err := btcjson.NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (cmd NameListCmd) UnmarshalJSON(b []byte) error {
	var r btcjson.RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 1 {
		return btcjson.ErrWrongNumberOfParams
	}

	newCmd, err := NameListFromRaw(&r)
	if err != nil {
		return err
	}

	cmd = newCmd.(NameListCmd)
	return nil
}

// NameHistoryCmd is a type handling custom marshaling and
// unmarshaling of name_history JSON RPC commands.
type NameHistoryCmd struct {
	id   interface{}
	Name string
}

// Enforce that NameHistoryCmd satisfies the Cmd interface.
var _ btcjson.Cmd = &NameHistoryCmd{}

// NewNameHistoryCmd creates a new NameHistoryCmd.
func NewNameHistoryCmd(id interface{}, name string) (*NameHistoryCmd, error) {
	return &NameHistoryCmd{
		id:   id,
		Name: name,
	}, nil
}

// NameHistoryFromRaw is a RawCmdParser.
func NameHistoryFromRaw(rawCmd *btcjson.RawCmd) (btcjson.Cmd, error) {
	var nameStr string
	err := json.Unmarshal(rawCmd.Params[0], &nameStr)
	if err != nil {
		return NameHistoryCmd{}, err
	}

	newCmd, err := NewNameHistoryCmd(rawCmd.Id, nameStr)
	if err != nil {
		return *newCmd, err
	}
	return *newCmd, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (cmd NameHistoryCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd NameHistoryCmd) Method() string {
	return "name_history"
}

// MarshalJSON returns the JSON encoding of cmd. Part of the Cmd interface.
func (cmd NameHistoryCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		cmd.Name,
	}

	raw, err := btcjson.NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (cmd NameHistoryCmd) UnmarshalJSON(b []byte) error {
	var r btcjson.RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return btcjson.ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("name parameter must be a string: %v", err)
	}

	newCmd, err := NewNameHistoryCmd(r.Id, name)
	if err != nil {
		return err
	}
	cmd = *newCmd
	return nil
}

// NameScanCmd is a type handling custom marshaling and
// unmarshaling of name_scan JSON RPC commands.
type NameScanCmd struct {
	id          interface{}
	StartName   string
	MaxReturned int
}

// Enforce that NameScanCmd satisfies the Cmd interface.
var _ btcjson.Cmd = &NameScanCmd{}

// NewNameScanCmd creates a new NameScanCmd.
func NewNameScanCmd(id interface{}, optArgs ...interface{}) (*NameScanCmd, error) {
	var startName string
	var maxReturned int

	if len(optArgs) > 2 {
		return nil, btcjson.ErrWrongNumberOfParams
	}
	if len(optArgs) > 0 {
		a, ok := optArgs[0].(string)
		if !ok {
			return nil, errors.New("first optional argument start-name is not a string")
		}
		startName = a
	}
	if len(optArgs) > 1 {
		b, ok := optArgs[1].(int)
		if !ok {
			return nil, errors.New("second optional argument max-returned is not an int")
		}
		maxReturned = b
	}

	return &NameScanCmd{
		id:          id,
		StartName:   startName,
		MaxReturned: maxReturned,
	}, nil
}

// NameScanFromRaw is a RawCmdParser.
func NameScanFromRaw(rawCmd *btcjson.RawCmd) (btcjson.Cmd, error) {
	var startName string
	var maxReturned int
	params := make([]interface{}, 0, 2)

	if len(rawCmd.Params) > 0 {
		err := json.Unmarshal(rawCmd.Params[0], &startName)
		if err != nil {
			return NameScanCmd{}, err
		}
		params = append(params, startName)
	}

	if len(rawCmd.Params) > 1 {
		err := json.Unmarshal(rawCmd.Params[1], &maxReturned)
		if err != nil {
			return NameScanCmd{}, err
		}
		params = append(params, maxReturned)
	}

	newCmd, err := NewNameScanCmd(rawCmd.Id, params)
	if err != nil {
		return NameScanCmd{}, err
	}
	return *newCmd, nil
}

// Id satisfies the Cmd interace by returning the id of the command.
func (cmd NameScanCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd NameScanCmd) Method() string {
	return "name_scan"
}

// MarshalJSON returns the JSON encoding of cmd. Part of the Cmd interface.
func (cmd NameScanCmd) MarshalJSON() ([]byte, error) {
	params := make([]interface{}, 0, 2)
	if cmd.StartName != "" {
		params = append(params, cmd.StartName)
	}
	if cmd.MaxReturned != 0 {
		params = append(params, cmd.MaxReturned)
	}

	raw, err := btcjson.NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (cmd NameScanCmd) UnmarshalJSON(b []byte) error {
	var r btcjson.RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 2 {
		return btcjson.ErrWrongNumberOfParams
	}

	newCmd, err := NameScanFromRaw(&r)
	if err != nil {
		return err
	}
	cmd = newCmd.(NameScanCmd)
	return nil
}

// NameFilterCmd is a type handling custom marshaling and
// unmarshaling of name_filter JSON RPC commands.
type NameFilterCmd struct {
	id     interface{}
	Regexp string
	MaxAge int
	From   int
	Nb     int
	Stat   int
}

// Enforce that NameFilterCmd satisfies the Cmd interface.
var _ btcjson.Cmd = &NameFilterCmd{}

// NewNameFilterCmd creates a new NameFilterCmd.
func NewNameFilterCmd(id interface{}, optArgs ...interface{}) (*NameFilterCmd, error) {
	if len(optArgs) > 5 {
		return nil, btcjson.ErrWrongNumberOfParams
	}
	var regexp string
	var maxage int
	var from int
	var nb int
	var stat int
	if len(optArgs) > 0 {
		a, ok := optArgs[0].(string)
		if !ok {
			return nil, errors.New("first optional argument regexp is not a string")
		}
		regexp = a
	}
	if len(optArgs) > 1 {
		b, ok := optArgs[1].(int)
		if !ok {
			return nil, errors.New("second optional argument maxage is not an int")
		}
		maxage = b
	}
	if len(optArgs) > 2 {
		c, ok := optArgs[2].(int)
		if !ok {
			return nil, errors.New("third optional argument from is not an int")
		}
		from = c
	}
	if len(optArgs) > 3 {
		d, ok := optArgs[3].(int)
		if !ok {
			return nil, errors.New("fourth optional argument nb is not an int")
		}
		nb = d
	}
	if len(optArgs) > 4 {
		e, ok := optArgs[4].(int)
		if !ok {
			return nil, errors.New("fifth optional argument stat is not an int")
		}
		stat = e
	}

	return &NameFilterCmd{
		id:     id,
		Regexp: regexp,
		MaxAge: maxage,
		From:   from,
		Nb:     nb,
		Stat:   stat,
	}, nil

}

// NameFilterFromRaw is a RawCmdParser.
func NameFilterFromRaw(rawCmd *btcjson.RawCmd) (btcjson.Cmd, error) {
	var regexp string
	var maxage int
	var from int
	var nb int
	var stat int
	params := make([]interface{}, 0, 5)

	if len(rawCmd.Params) > 0 {
		err := json.Unmarshal(rawCmd.Params[0], &regexp)
		if err != nil {
			return NameFilterCmd{}, err
		}
		params = append(params, regexp)
	}
	if len(rawCmd.Params) > 1 {
		err := json.Unmarshal(rawCmd.Params[1], &maxage)
		if err != nil {
			return NameFilterCmd{}, err
		}
		params = append(params, maxage)
	}
	if len(rawCmd.Params) > 2 {
		err := json.Unmarshal(rawCmd.Params[2], &from)
		if err != nil {
			return NameFilterCmd{}, err
		}
		params = append(params, from)
	}
	if len(rawCmd.Params) > 3 {
		err := json.Unmarshal(rawCmd.Params[3], &nb)
		if err != nil {
			return NameFilterCmd{}, err
		}
		params = append(params, nb)
	}
	if len(rawCmd.Params) > 4 {
		err := json.Unmarshal(rawCmd.Params[4], &stat)
		if err != nil {
			return NameFilterCmd{}, err
		}
		params = append(params, stat)
	}

	newCmd, err := NewNameFilterCmd(rawCmd.Id, params)
	if err != nil {
		return NameFilterCmd{}, err
	}
	return *newCmd, nil
}

// Id satisfies the Cmd interace by returning the id of the command.
func (cmd NameFilterCmd) Id() interface{} {
	return cmd.id
}

// Method satisfies the Cmd interface by returning the json method.
func (cmd NameFilterCmd) Method() string {
	return "name_filter"
}

// MarshalJSON returns the JSON encoding of cmd. Part of the Cmd interface.
func (cmd NameFilterCmd) MarshalJSON() ([]byte, error) {
	params := make([]interface{}, 0, 5)
	if cmd.Regexp != "" {
		params = append(params, cmd.Regexp)
	}
	if cmd.MaxAge != 0 {
		params = append(params, cmd.MaxAge)
	}
	if cmd.From != 0 {
		params = append(params, cmd.From)
	}
	if cmd.Nb != 0 {
		params = append(params, cmd.Nb)
	}
	if cmd.Stat != 0 {
		params = append(params, cmd.Stat)
	}

	raw, err := btcjson.NewRawCmd(cmd.id, cmd.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (cmd NameFilterCmd) UnmarshalJSON(b []byte) error {
	var r btcjson.RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 5 {
		return btcjson.ErrWrongNumberOfParams
	}

	newCmd, err := NameFilterFromRaw(&r)
	if err != nil {
		return err
	}
	cmd = newCmd.(NameFilterCmd)
	return nil
}
