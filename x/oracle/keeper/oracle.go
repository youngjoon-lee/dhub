package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func (k Keeper) SetOracle(ctx sdk.Context, oracle types.Oracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OracleKeyPrefix)
	b := k.cdc.MustMarshal(&oracle)
	store.Set(types.OracleKey(oracle.OperatorAddress), b)
}

func (k Keeper) GetOracle(ctx sdk.Context, operatorAddress string) (oracle types.Oracle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OracleKeyPrefix)

	b := store.Get(types.OracleKey(operatorAddress))
	if b == nil {
		return oracle, false
	}

	k.cdc.MustUnmarshal(b, &oracle)
	return oracle, true
}

func (k Keeper) GetAllOracles(ctx sdk.Context) (list []types.Oracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OracleKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oracle types.Oracle
		k.cdc.MustUnmarshal(iterator.Value(), &oracle)
		list = append(list, oracle)
	}

	return
}

func (k Keeper) GetOracleCount(ctx sdk.Context) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OracleKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	cnt := uint32(0)
	for ; iterator.Valid(); iterator.Next() {
		cnt++
	}
	return cnt
}
