package keeper

import (
	"fmt"
	"github.com/bcflapp/x/bcflapp/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*

In db, we store key & value.

key: ([]byte) as a job id
value: ([]byte) job's information

 */

func (k Keeper) CreateApply(ctx sdk.Context, apply types.Apply) {
	//var joblistList []types.Joblist
	//store := ctx.KVStore(k.storeKey)
	//iterator := sdk.KVStorePrefixIterator(store, []byte(types.JoblistPrefix))
	//var kklist []string
	//for ; iterator.Valid(); iterator.Next() {
	//	var joblist types.Joblist
	//	k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &joblist)
	//	kklist = append(kklist, string(iterator.Key()))
	//	joblistList = append(joblistList, joblist)
	//}
	/////////////////////////////////////
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.JoblistPrefix + apply.Jobid)
	value := store.Get(key)
	if value == nil {
		fmt.Sprintf("No such job.")
		//fmt.Errorf("failed")
		//k.Logger2(ctx, "No such job")

		//k.Logger(ctx)
		panic("No such job.")
		//fmt.Fprintln("No such job.")
		//return nil, sdk.Result{}
	}
	//store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(new(types.Joblist)))
	//fmt.Printf("Value in keeper: %s\n\n", value)
	//
	var job types.Joblist
	//MustUnmarshalBinaryLengthPrefixed
	//k.cdc.MustUnmarshalBinaryLengthPrefixed(value, job)
	//err := k.cdc.UnmarshalJSON(value, job)
	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &job)
	//
	//if err != nil {
	//	//panic(fmt.Sprintf("Invalid job stored: %s", err))
	//}
	//
	////apply.Creator
	NewMember := types.MemberState{
		Member:      apply.Creator,
		Tags:        apply.Tags,
		Datapath:    apply.Datapath,
		Auth:        false,
		TrainJob:    "",
		TrainState:  -1,
		TrainResult: "",
	}

	//
	//
	////	MsgCreateApply{
	////	ID:       apply.ID,
	////	Creator:  apply.Creator,
	////	Jobid:    types.ApplyPrefix + apply.ID,
	////	Tags:     apply.Tags,
	////	Datapath: apply.Datapath,
	////}
	//
	job.Members = append(job.Members, NewMember)
	//value := k.cdc.MustMarshalBinaryLengthPrefixed()
	NewValue := k.cdc.MustMarshalBinaryLengthPrefixed(job)
	store.Set(key, NewValue)

	//store := ctx.KVStore(k.storeKey)
	//key := []byte(types.JoblistPrefix + apply.ID)
	//value := k.cdc.MustMarshalBinaryLengthPrefixed(apply)
	//store.Set(key, value)

	//return  job, sdk.Result{}
}

func listApply(ctx sdk.Context, k Keeper) ([]byte, error) {
  var applyList []types.Apply
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.ApplyPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var apply types.Apply
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &apply)
    applyList = append(applyList, apply)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, applyList)
  return res, nil
}