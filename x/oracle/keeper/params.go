package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MaxOracles(ctx),
		k.SlashFractionDowntime(ctx),
		k.SlashFractionWrongVote(ctx),
		k.DowntimeDuration(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MaxOracles returns the MaxOracles param
func (k Keeper) MaxOracles(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxOracles, &res)
	return
}

// SlashFractionDowntime returns the SlashFractionDowntime param
func (k Keeper) SlashFractionDowntime(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeySlashFractionDowntime, &res)
	return
}

// SlashFractionWrongVote returns the SlashFractionWrongVote param
func (k Keeper) SlashFractionWrongVote(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeySlashFractionWrongVote, &res)
	return
}

// DowntimeDuration returns the DowntimeDuration param
func (k Keeper) DowntimeDuration(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyDowntimeDuration, &res)
	return
}
