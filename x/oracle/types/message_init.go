package types

import (
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInit = "init"

var _ sdk.Msg = &MsgInit{}

func NewMsgInit(operatorAddress string, enclaveReport []byte, oraclePubKey *secp256k1.PubKey) *MsgInit {
	return &MsgInit{
		OperatorAddress: operatorAddress,
		EnclaveReport:   enclaveReport,
		OraclePubKey:    oraclePubKey,
	}
}

func (msg *MsgInit) Route() string {
	return RouterKey
}

func (msg *MsgInit) Type() string {
	return TypeMsgInit
}

func (msg *MsgInit) GetSigners() []sdk.AccAddress {
	operatorAddress, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operatorAddress}
}

func (msg *MsgInit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operatorAddress address (%s)", err)
	}

	if len(msg.EnclaveReport) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "enclaveReport is empty")
	}

	if len(msg.OraclePubKey.Key) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "oraclePubKey is empty")
	}

	return nil
}
