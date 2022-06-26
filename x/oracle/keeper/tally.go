package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func (k Keeper) Tally(ctx sdk.Context, join types.Join) (done bool, approved bool, tallyResults types.TallyResult) {
	tallyResult := types.TallyResult{
		Yes:      sdk.ZeroInt(),
		YesValue: "",
		No:       sdk.ZeroInt(),
	}
	yesValues := make(map[string]sdk.Int, 0)

	//TODO: get the total voting power of existing oracle nodes
	totalVotingPower := sdk.OneInt()

	k.IterateVoteForJoins(ctx, join.ID, func(vote types.VoteForJoin) bool {
		//TODO: adopt voting powers
		switch vote.Option {
		case types.OptionYes:
			tallyResult.Yes = tallyResult.Yes.Add(sdk.OneInt())
			addYesValue(yesValues, vote)
		case types.OptionNo:
			tallyResult.No = tallyResult.No.Add(sdk.OneInt())
		}

		return false
	})

	half := totalVotingPower.Quo(sdk.NewInt(2))

	if tallyResult.No.GT(half) {
		return true, false, tallyResult
	}
	if tallyResult.Yes.GT(half) {
		yesValue, ok := getTopYesValue(yesValues)
		if !ok {
			return false, false, tallyResult
		}
		tallyResult.YesValue = yesValue
		return true, true, tallyResult
	}

	return false, false, tallyResult
}

func addYesValue(yesValues map[string]sdk.Int, vote types.VoteForJoin) {
	//TODO: adopt voting powers
	value, ok := yesValues[vote.EncryptedOraclePrivKeyB64]
	if !ok {
		yesValues[vote.EncryptedOraclePrivKeyB64] = sdk.OneInt()
	} else {
		yesValues[vote.EncryptedOraclePrivKeyB64] = value.Add(sdk.OneInt())
	}
}

func getTopYesValue(yesValues map[string]sdk.Int) (string, bool) {
	topValues := make([]string, 0, 1)
	topPower := sdk.ZeroInt()

	for value, power := range yesValues {
		if power.GT(topPower) {
			topValues = []string{value}
		} else if power.Equal(topPower) {
			topValues = append(topValues, value)
		}
	}

	if len(topValues) != 1 {
		return "", false
	}
	return topValues[0], true
}
