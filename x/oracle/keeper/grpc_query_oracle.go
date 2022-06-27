package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OracleAll(c context.Context, req *types.QueryAllOracleRequest) (*types.QueryAllOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var oracles []types.Oracle
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	oracleStore := prefix.NewStore(store, types.OracleKeyPrefix)

	pageRes, err := query.Paginate(oracleStore, req.Pagination, func(key []byte, value []byte) error {
		var oracle types.Oracle
		if err := k.cdc.Unmarshal(value, &oracle); err != nil {
			return err
		}

		oracles = append(oracles, oracle)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOracleResponse{Oracle: oracles, Pagination: pageRes}, nil
}

func (k Keeper) Oracle(c context.Context, req *types.QueryGetOracleRequest) (*types.QueryGetOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOracle(ctx, req.OperatorAddress)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOracleResponse{Oracle: val}, nil
}
