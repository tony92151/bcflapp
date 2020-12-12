package bcflapp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bcflapp/x/bcflapp/keeper"
	"github.com/bcflapp/x/bcflapp/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreateApply:
			return handleMsgCreateApply(ctx, k, msg)
		case types.MsgCreateJoblist:
			return handleMsgCreateJoblist(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
