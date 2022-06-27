package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextJoinID: 1,
		JoinList:   []Join{},
		OracleList: []Oracle{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in oracle
	oracleKeys := make(map[string]struct{})

	for _, oracle := range gs.OracleList {
		keyStr := string(OracleKey(oracle.OperatorAddress))
		if _, ok := oracleKeys[keyStr]; ok {
			return fmt.Errorf("duplicated id for oracle")
		}
		oracleKeys[keyStr] = struct{}{}
	}

	// Check for duplicated index in join
	joinKeys := make(map[string]struct{})

	for _, join := range gs.JoinList {
		keyStr := string(JoinKey(join.ID))
		if _, ok := joinKeys[keyStr]; ok {
			return fmt.Errorf("duplicated id for join")
		}
		joinKeys[keyStr] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
