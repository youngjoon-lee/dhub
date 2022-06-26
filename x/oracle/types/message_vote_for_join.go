package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteForJoin = "voteForJoin"

var _ sdk.Msg = &MsgVoteForJoin{}

func NewMsgVoteForJoin(joinID uint64, option VoteOption, encryptedOraclePrivKeyB64 string, voter string) *MsgVoteForJoin {
	return &MsgVoteForJoin{
		JoinID:                    joinID,
		Option:                    option,
		EncryptedOraclePrivKeyB64: encryptedOraclePrivKeyB64,
		Voter:                     voter,
	}
}

func (msg *MsgVoteForJoin) Route() string {
	return RouterKey
}

func (msg *MsgVoteForJoin) Type() string {
	return TypeMsgVoteForJoin
}

func (msg *MsgVoteForJoin) GetSigners() []sdk.AccAddress {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voter}
}

func (msg *MsgVoteForJoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteForJoin) ValidateBasic() error {
	if len(msg.EncryptedOraclePrivKeyB64) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "encryptedOraclePrivKey is empty")
	}
	if _, err := sdk.AccAddressFromBech32(msg.Voter); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operatorAddress address (%s)", err)
	}
	return nil
}
