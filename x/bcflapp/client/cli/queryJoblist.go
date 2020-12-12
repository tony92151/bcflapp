package cli

import (
	"fmt"
	"github.com/bcflapp/x/bcflapp/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetCmdListJoblist(queryRoute string, cdc *codec.Codec) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list-joblist",
		Short: "list all joblist",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListJoblist, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Joblist\n%s\n", err.Error())
				return nil
			}
			var out []types.Joblist
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
	listCmd.Flags().Bool("hire", true, "Only list hiring jobs.")
	return listCmd
}

//

func GetCmdListJoblistFromID(queryRoute string, cdc *codec.Codec) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list-joblist-id",
		Short: "list job from id",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListJoblistId, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Joblist\n%s\n", err.Error())
				return nil
			}
			var out []types.Joblist
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
	listCmd.Flags().Bool("hire", true, "Only list hiring jobs.")
	return listCmd
}


