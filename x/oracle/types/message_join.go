package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJoin = "join"

var _ sdk.Msg = &MsgJoin{}

func NewMsgJoin(operatorAddress string, enclaveReport string) *MsgJoin {
	return &MsgJoin{
		OperatorAddress: operatorAddress,
		EnclaveReport:   enclaveReport,
	}
}

func (msg *MsgJoin) Route() string {
	return RouterKey
}

func (msg *MsgJoin) Type() string {
	return TypeMsgJoin
}

func (msg *MsgJoin) GetSigners() []sdk.AccAddress {
	operatorAddress, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operatorAddress}
}

func (msg *MsgJoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operatorAddress address (%s)", err)
	}
	return nil
}
