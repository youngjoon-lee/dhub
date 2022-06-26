package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultTallyResult() TallyResult {
	return TallyResult{
		Yes:      sdk.ZeroInt(),
		YesValue: "",
		No:       sdk.ZeroInt(),
	}
}
