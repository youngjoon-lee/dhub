package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	ErrUnknownJoin  = sdkerrors.Register(ModuleName, 2, "unknown join")
	ErrClosedJoin   = sdkerrors.Register(ModuleName, 3, "join was already closed")
	ErrJoinExists   = sdkerrors.Register(ModuleName, 4, "join already exists")
	ErrOracleExists = sdkerrors.Register(ModuleName, 5, "oracle already exists")
)
