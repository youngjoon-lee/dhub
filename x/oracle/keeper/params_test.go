package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/youngjoon-lee/dhub/testutil/keeper"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MaxOracles, k.MaxOracles(ctx))
	require.EqualValues(t, params.SlashFractionDowntime, k.SlashFractionDowntime(ctx))
	require.EqualValues(t, params.SlashFractionWrongVote, k.SlashFractionWrongVote(ctx))
	require.EqualValues(t, params.DowntimeDuration, k.DowntimeDuration(ctx))
}
