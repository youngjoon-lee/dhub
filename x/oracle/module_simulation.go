package oracle

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/youngjoon-lee/dhub/testutil/sample"
	oraclesimulation "github.com/youngjoon-lee/dhub/x/oracle/simulation"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oraclesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgJoin = "op_weight_msg_join"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJoin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oracleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oracleGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	oracleParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxOracles), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(oracleParams.MaxOracles))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeySlashFractionDowntime), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(oracleParams.SlashFractionDowntime))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeySlashFractionWrongVote), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(oracleParams.SlashFractionWrongVote))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyDowntimeDuration), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(oracleParams.DowntimeDuration))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgJoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgJoin, &weightMsgJoin, nil,
		func(_ *rand.Rand) {
			weightMsgJoin = defaultWeightMsgJoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoin,
		oraclesimulation.SimulateMsgJoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
