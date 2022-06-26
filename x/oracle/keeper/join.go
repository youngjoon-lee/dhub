package keeper

import (
	"encoding/base64"
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// SetJoin set a specific join in the store from its index
func (k Keeper) SetJoin(ctx sdk.Context, join types.Join) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JoinKeyPrefix)
	b := k.cdc.MustMarshal(&join)
	store.Set(types.JoinKey(join.ID), b)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeJoin,
			sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(join.ID, 10)),
			sdk.NewAttribute(types.AttributeKeyOperatorAddress, join.OperatorAddress),
			sdk.NewAttribute(types.AttributeKeyEnclaveReportBase64, base64.StdEncoding.EncodeToString(join.EnclaveReport)),
			sdk.NewAttribute(types.AttributeKeyEncPubKeyBase64, base64.StdEncoding.EncodeToString(join.EncPubKey.Key)),
		),
	)
}

// GetJoin returns a join from its index
func (k Keeper) GetJoin(ctx sdk.Context, joinID uint64) (val types.Join, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JoinKeyPrefix)

	b := store.Get(types.JoinKey(joinID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllJoin returns all join
func (k Keeper) GetAllJoin(ctx sdk.Context) (list []types.Join) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JoinKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Join
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetNextJoinID(ctx sdk.Context, id uint64) {
	idBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(idBytes, id)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextJoinIDKey, idBytes)
}

func (k Keeper) GetNextJoinID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	idBytes := store.Get(types.NextJoinIDKey)
	if idBytes == nil {
		panic("NextJoinIDKey is not found")
	}
	return binary.LittleEndian.Uint64(idBytes)
}
