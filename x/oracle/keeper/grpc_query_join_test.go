package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/youngjoon-lee/dhub/testutil/keeper"
	"github.com/youngjoon-lee/dhub/testutil/nullify"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestJoinQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNJoin(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetJoinRequest
		response *types.QueryGetJoinResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetJoinRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetJoinResponse{Join: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetJoinRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetJoinResponse{Join: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetJoinRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Join(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestJoinQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNJoin(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllJoinRequest {
		return &types.QueryAllJoinRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.JoinAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Join), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Join),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.JoinAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Join), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Join),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.JoinAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Join),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.JoinAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
