package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/youngjoon-lee/dhub/testutil/keeper"
	"github.com/youngjoon-lee/dhub/testutil/nullify"
	"github.com/youngjoon-lee/dhub/x/oracle/keeper"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNJoin(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Join {
	items := make([]types.Join, n)
	for i := range items {
		items[i].JoinId = uint64(i)

		keeper.SetJoin(ctx, items[i])
	}
	return items
}

func TestJoinGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNJoin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetJoin(ctx,
			item.JoinId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestJoinRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNJoin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveJoin(ctx,
			item.JoinId,
		)
		_, found := keeper.GetJoin(ctx,
			item.JoinId,
		)
		require.False(t, found)
	}
}

func TestJoinGetAll(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNJoin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllJoin(ctx)),
	)
}
