package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InsertToPendingJoinQueue(ctx sdk.Context, joinID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingJoinQueueKeyPrefix)
	store.Set(types.PendingJoinQueueKey(joinID), types.GetJoinIDBytes(joinID))
}

func (k Keeper) RemoveFromPendingJoinQueue(ctx sdk.Context, joinID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingJoinQueueKeyPrefix)
	store.Delete(types.PendingJoinQueueKey(joinID))
}

func (k Keeper) IteratePendingJoinQueue(ctx sdk.Context, cb func(join types.Join) (stop bool)) {
	iterator := k.PendingJoinQueueIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		joinID := types.SplitPendingJoinQueueKey(iterator.Key())
		join, ok := k.GetJoin(ctx, joinID)
		if !ok {
			panic(fmt.Sprintf("join %d not found", joinID))
		}

		if stop := cb(join); stop {
			break
		}
	}
}

func (k Keeper) PendingJoinQueueIterator(ctx sdk.Context) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PendingJoinQueueKeyPrefix)
	return store.Iterator(nil, nil)
}
