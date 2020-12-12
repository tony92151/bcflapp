package keeper

import (
  // this line is used by starport scaffolding
	"github.com/bcflapp/x/bcflapp/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for bcflapp clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListApply:
			return listApply(ctx, k)
		case types.QueryListJoblist:
			return listJoblist(ctx, k)
		case types.QueryListJoblistId:
			return listJoblistid(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown bcflapp query endpoint")
		}
	}
}