package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxOracles = []byte("MaxOracles")
	// TODO: Determine the default value
	DefaultMaxOracles uint64 = 0
)

var (
	KeySlashFractionDowntime = []byte("SlashFractionDowntime")
	// TODO: Determine the default value
	DefaultSlashFractionDowntime string = "slash_fraction_downtime"
)

var (
	KeySlashFractionWrongVote = []byte("SlashFractionWrongVote")
	// TODO: Determine the default value
	DefaultSlashFractionWrongVote string = "slash_fraction_wrong_vote"
)

var (
	KeyDowntimeDuration = []byte("DowntimeDuration")
	// TODO: Determine the default value
	DefaultDowntimeDuration string = "downtime_duration"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxOracles uint64,
	slashFractionDowntime string,
	slashFractionWrongVote string,
	downtimeDuration string,
) Params {
	return Params{
		MaxOracles:             maxOracles,
		SlashFractionDowntime:  slashFractionDowntime,
		SlashFractionWrongVote: slashFractionWrongVote,
		DowntimeDuration:       downtimeDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxOracles,
		DefaultSlashFractionDowntime,
		DefaultSlashFractionWrongVote,
		DefaultDowntimeDuration,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxOracles, &p.MaxOracles, validateMaxOracles),
		paramtypes.NewParamSetPair(KeySlashFractionDowntime, &p.SlashFractionDowntime, validateSlashFractionDowntime),
		paramtypes.NewParamSetPair(KeySlashFractionWrongVote, &p.SlashFractionWrongVote, validateSlashFractionWrongVote),
		paramtypes.NewParamSetPair(KeyDowntimeDuration, &p.DowntimeDuration, validateDowntimeDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxOracles(p.MaxOracles); err != nil {
		return err
	}

	if err := validateSlashFractionDowntime(p.SlashFractionDowntime); err != nil {
		return err
	}

	if err := validateSlashFractionWrongVote(p.SlashFractionWrongVote); err != nil {
		return err
	}

	if err := validateDowntimeDuration(p.DowntimeDuration); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMaxOracles validates the MaxOracles param
func validateMaxOracles(v interface{}) error {
	maxOracles, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxOracles

	return nil
}

// validateSlashFractionDowntime validates the SlashFractionDowntime param
func validateSlashFractionDowntime(v interface{}) error {
	slashFractionDowntime, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = slashFractionDowntime

	return nil
}

// validateSlashFractionWrongVote validates the SlashFractionWrongVote param
func validateSlashFractionWrongVote(v interface{}) error {
	slashFractionWrongVote, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = slashFractionWrongVote

	return nil
}

// validateDowntimeDuration validates the DowntimeDuration param
func validateDowntimeDuration(v interface{}) error {
	downtimeDuration, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = downtimeDuration

	return nil
}
