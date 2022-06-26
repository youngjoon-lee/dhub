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
		EncPubKey:       msg.EncPubKey,
		Status:          types.JOIN_STATUS_PENDING,
		TallyResult:     types.DefaultTallyResult(),
	}
	k.SetJoin(ctx, join)
	k.InsertToPendingJoinQueue(ctx, join.ID)
	k.SetNextJoinID(ctx, join.ID+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
		),
	)

	return &types.MsgJoinResponse{ID: join.ID}, nil
}
