package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateJoblist{}

type MsgCreateJoblist struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Jobcreator string `json:"jobcreator" yaml:"jobcreator"`
  Tags []string `json:"tags" yaml:"tags"`
  Limit string `json:"limit" yaml:"limit"`
  Members []MemberState `json:"members" yaml:"members"`
}

//type MemberState struct {
//  Member sdk.AccAddress
//  Auth bool
//  TrainJob string
//  TrainState int  // -1: waiting, 0: authorized and get job info, 1: training, 2: done
//  TrainResult string
//}

func NewMsgCreateJoblist(creator sdk.AccAddress, jobcreator string, tags []string, limit string, members []MemberState) MsgCreateJoblist {
  return MsgCreateJoblist{
    ID: uuid.New().String(),
		Creator: creator,
    Jobcreator: jobcreator,
    Tags: tags,
    Limit: limit,
    Members: members,
	}
}

func (msg MsgCreateJoblist) Route() string {
  return RouterKey
}

func (msg MsgCreateJoblist) Type() string {
  return "CreateJoblist"
}

func (msg MsgCreateJoblist) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateJoblist) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateJoblist) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}