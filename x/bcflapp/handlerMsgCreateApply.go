package bcflapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bcflapp/x/bcflapp/types"
	"github.com/bcflapp/x/bcflapp/keeper"
)

func handleMsgCreateApply(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateApply) (*sdk.Result, error) {
	var apply = types.Apply{
		Creator: msg.Creator,
		ID:      msg.ID,
    Jobid: msg.Jobid,
    Tags: msg.Tags,
    Datapath: msg.Datapath,
	}
	k.CreateApply(ctx, apply)

	//if job == nil{
	//	return &res, nil
	//}


	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
