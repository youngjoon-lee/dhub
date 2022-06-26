package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// AddVoteForJoin adds a vote on a specific join
func (k Keeper) SetVoteForJoin(ctx sdk.Context, vote types.VoteForJoin) error {
	join, ok := k.GetJoin(ctx, vote.JoinID)
	if !ok {
		return sdkerrors.Wrapf(types.ErrUnknownJoin, "%v", vote.JoinID)
	}
	if join.Status != types.JOIN_STATUS_PENDING {
		return sdkerrors.Wrapf(types.ErrClosedJoin, "%v", vote.JoinID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VoteForJoinKeyPrefix)
	b := k.cdc.MustMarshal(&vote)
	store.Set(types.VoteForJoinKey(vote.JoinID, vote.Voter), b)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVoteForJoin,
			sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(vote.JoinID, 10)),
			sdk.NewAttribute(types.AttributeKeyOption, vote.Option.String()),
			sdk.NewAttribute(types.AttributeKeyVoter, vote.Voter),
		),
	)

	return nil
}

func (k Keeper) IterateVoteForJoins(ctx sdk.Context, joinID uint64, cb func(vote types.VoteForJoin) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.VoteForJoinKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, types.VoteForJoinsKey(joinID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var vote types.VoteForJoin
		k.cdc.MustUnmarshal(iterator.Value(), &vote)

		if stop := cb(vote); stop {
			break
		}
	}

}
