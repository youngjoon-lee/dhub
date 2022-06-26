package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func (k msgServer) VoteForJoin(goCtx context.Context, msg *types.MsgVoteForJoin) (*types.MsgVoteForJoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	vote := types.VoteForJoin{
		JoinID:                    msg.JoinID,
		Option:                    msg.Option,
		EncryptedOraclePrivKeyB64: msg.EncryptedOraclePrivKeyB64,
		Voter:                     msg.Voter,
	}

	if err := k.SetVoteForJoin(ctx, vote); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
		),
	)

	return &types.MsgVoteForJoinResponse{}, nil
}
