package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracleCnt := k.GetOracleCount(ctx)
	if oracleCnt != 0 {
		return nil, types.ErrOraclesExist
	}

	oracle := types.Oracle{
		OperatorAddress: msg.OperatorAddress,
		EnclaveReport:   msg.EnclaveReport,
		Stake:           sdk.OneInt(), //TODO: proof-of-stake
	}

	k.SetOracle(ctx, oracle)
	k.SetOraclePubKey(ctx, types.OraclePubKey{PubKey: msg.OraclePubKey})

	return &types.MsgInitResponse{}, nil
}
