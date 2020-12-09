package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Joblist struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Jobcreator string `json:"jobcreator" yaml:"jobcreator"`
  Tags []string `json:"tags" yaml:"tags"`
  Limit string `json:"limit" yaml:"limit"`
  Members []string `json:"members" yaml:"members"`
}