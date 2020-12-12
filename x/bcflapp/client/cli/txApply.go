package cli

import (
	"bufio"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/bcflapp/x/bcflapp/types"
)

func GetCmdCreateApply(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-apply [jobid] [tags] [datapath]",
		Short: "Creates a new apply",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsJobid := string(args[0])
      //argsTags := string(args[1])
      argsTags := strings.Split(string(args[1]), ",")
      argsDatapath := string(args[2])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateApply(cliCtx.GetFromAddress(), argsJobid, argsTags, argsDatapath)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}