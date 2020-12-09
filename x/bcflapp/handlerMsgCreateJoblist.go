package bcflapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bcflapp/x/bcflapp/types"
	"github.com/bcflapp/x/bcflapp/keeper"
)

func handleMsgCreateJoblist(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateJoblist) (*sdk.Result, error) {
	var joblist = types.Joblist{
		Creator: msg.Creator,
		ID:      msg.ID,
    Jobcreator: msg.Jobcreator,
    Tags: msg.Tags,
    Limit: msg.Limit,
    Members: msg.Members,
	}
	k.CreateJoblist(ctx, joblist)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
