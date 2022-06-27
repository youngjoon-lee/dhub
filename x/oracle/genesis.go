package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/keeper"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the oracle
	for _, elem := range genState.OracleList {
		k.SetOracle(ctx, elem)
	}

	k.SetNextJoinID(ctx, genState.NextJoinID)
	// Set all the join
	for _, elem := range genState.JoinList {
		k.SetJoin(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.OracleList = k.GetAllOracles(ctx)
	genesis.NextJoinID = k.GetNextJoinID(ctx)
	genesis.JoinList = k.GetAllJoin(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
