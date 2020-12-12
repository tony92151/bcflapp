package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Apply struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Jobid string `json:"jobid" yaml:"jobid"`
  Tags []string `json:"tags" yaml:"tags"`
  Datapath string `json:"datapath" yaml:"datapath"`
}