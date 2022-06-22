package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/youngjoon-lee/dhub/testutil/keeper"
	"github.com/youngjoon-lee/dhub/testutil/nullify"
	"github.com/youngjoon-lee/dhub/x/oracle"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
