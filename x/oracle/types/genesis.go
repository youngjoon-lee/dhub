package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		JoinList: []Join{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in join
	joinIDMap := make(map[string]struct{})

	for _, elem := range gs.JoinList {
		index := string(JoinKey(elem.JoinId))
		if _, ok := joinIDMap[index]; ok {
			return fmt.Errorf("duplicated id for join")
		}
		joinIDMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
