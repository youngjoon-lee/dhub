package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// SetJoin set a specific join in the store from its index
func (k Keeper) SetJoin(ctx sdk.Context, join types.Join) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JoinKeyPrefix))
	b := k.cdc.MustMarshal(&join)
	store.Set(types.JoinKey(
		join.JoinId,
	), b)
}

// GetJoin returns a join from its index
func (k Keeper) GetJoin(
	ctx sdk.Context,
	joinID uint64,

) (val types.Join, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JoinKeyPrefix))

	b := store.Get(types.JoinKey(
		joinID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveJoin removes a join from the store
func (k Keeper) RemoveJoin(
	ctx sdk.Context,
	joinID uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JoinKeyPrefix))
	store.Delete(types.JoinKey(
		joinID,
	))
}

// GetAllJoin returns all join
func (k Keeper) GetAllJoin(ctx sdk.Context) (list []types.Join) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.JoinKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Join
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
