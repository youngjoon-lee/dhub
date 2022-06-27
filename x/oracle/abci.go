package oracle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/keeper"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	logger := ctx.Logger()

	logger.Info("iterating pending join queue")
	keeper.IteratePendingJoinQueue(ctx, func(join types.Join) bool {
		done, approved, tallyResult := keeper.Tally(ctx, join)
		if !done {
			return false
		}
		logger.Info(fmt.Sprintf("joinID:%v, approved:%v, tally:%v", join.ID, approved, tallyResult))

		if approved {
			join.Status = types.JOIN_STATUS_APPROVED
		} else {
			join.Status = types.JOIN_STATUS_REJECTED
		}
		join.TallyResult = tallyResult
		keeper.SetJoin(ctx, join)

		oracle := types.Oracle{
			OperatorAddress: join.OperatorAddress,
			Stake:           sdk.OneInt(), //TODO: proof-of-stake
		}
		keeper.SetOracle(ctx, oracle)

		keeper.RemoveFromPendingJoinQueue(ctx, join.ID)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeJoinResult,
				sdk.NewAttribute(types.AttributeKeyID, strconv.FormatUint(join.ID, 10)),
				sdk.NewAttribute(types.AttributeKeyStatus, join.Status.String()),
				sdk.NewAttribute(types.AttributeKeyValue, join.TallyResult.YesValue),
			),
		)

		return false
	})
}
