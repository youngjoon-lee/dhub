package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/youngjoon-lee/dhub/x/oracle/keeper"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgInit:
			res, err := msgServer.Init(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgJoin:
			res, err := msgServer.Join(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgVoteForJoin:
			res, err := msgServer.VoteForJoin(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
