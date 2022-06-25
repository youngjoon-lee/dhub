package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func (k msgServer) Join(goCtx context.Context, msg *types.MsgJoin) (*types.MsgJoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	join := types.Join{
		ID:              k.GetNextJoinID(ctx),
		OperatorAddress: msg.OperatorAddress,
		EnclaveReport:   msg.EnclaveReport,
		Status:          types.JOIN_STATUS_PENDING,
	}
	k.SetJoin(ctx, join)
	k.SetNextJoinID(ctx, join.ID+1)

	return &types.MsgJoinResponse{ID: join.ID}, nil
}
