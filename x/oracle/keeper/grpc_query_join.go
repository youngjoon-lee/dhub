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

func (k Keeper) JoinAll(c context.Context, req *types.QueryAllJoinRequest) (*types.QueryAllJoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var joins []types.Join
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	joinStore := prefix.NewStore(store, types.KeyPrefix(types.JoinKeyPrefix))

	pageRes, err := query.Paginate(joinStore, req.Pagination, func(key []byte, value []byte) error {
		var join types.Join
		if err := k.cdc.Unmarshal(value, &join); err != nil {
			return err
		}

		joins = append(joins, join)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllJoinResponse{Join: joins, Pagination: pageRes}, nil
}

func (k Keeper) Join(c context.Context, req *types.QueryGetJoinRequest) (*types.QueryGetJoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetJoin(
		ctx,
		req.ID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetJoinResponse{Join: val}, nil
}
