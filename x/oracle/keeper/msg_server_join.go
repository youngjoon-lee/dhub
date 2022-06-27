package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if _, found := k.GetJoin(ctx, join.ID); found {
		return nil, sdkerrors.Wrapf(types.ErrJoinExists, "%v", join.ID)
	}
	if _, found := k.GetOracle(ctx, join.OperatorAddress); found {
		return nil, sdkerrors.Wrapf(types.ErrOracleExists, "%v", join.OperatorAddress)
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
