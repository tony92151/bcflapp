package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bcflapp/x/bcflapp/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateJoblist(ctx sdk.Context, joblist types.Joblist) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.JoblistPrefix + joblist.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(joblist)
	store.Set(key, value)
}

func listJoblist(ctx sdk.Context, k Keeper) ([]byte, error) {
  var joblistList []types.Joblist
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.JoblistPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var joblist types.Joblist
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &joblist)
    joblistList = append(joblistList, joblist)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, joblistList)
  return res, nil
}

func listJoblistid(ctx sdk.Context, k Keeper) ([]byte, error) {
	var joblistList []types.Joblist
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.JoblistPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var joblist types.Joblist
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &joblist)
		joblistList = append(joblistList, joblist)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, joblistList)
	return res, nil
}