package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateApply{}

type MsgCreateApply struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Jobid string `json:"jobid" yaml:"jobid"`
  Tags []string `json:"tags" yaml:"tags"`
  Datapath string `json:"datapath" yaml:"datapath"`
}

type MemberState struct {
	Member sdk.AccAddress
	Tags []string
	Datapath string
	Auth bool
	TrainJob string
	TrainState int  // -1: waiting, 0: authorized and get job info, 1: training, 2: done
	TrainResult string
}

func NewMsgCreateApply(creator sdk.AccAddress, jobid string, tags []string, dbpath string) MsgCreateApply {
  return MsgCreateApply{
    ID:          uuid.New().String(),
    Creator: creator,
    Jobid:       jobid,
    Tags:        tags,
    Datapath:    dbpath,
	}
}

func (msg MsgCreateApply) Route() string {
  return RouterKey
}

func (msg MsgCreateApply) Type() string {
  return "CreateApply"
}

func (msg MsgCreateApply) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateApply) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateApply) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}