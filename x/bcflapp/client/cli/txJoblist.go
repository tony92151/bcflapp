package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/bcflapp/x/bcflapp/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetCmdCreateJoblist(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-joblist [jobcreator] [tags] [limit] [members]",
		Short: "Creates a new joblist",
		Long : "create-joblist cosmos... price,size 3 cosmos...,cosmos...",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsJobcreator := string(args[0])
			argsTags := strings.Split((string(args[1])), ",")
			argsLimit := string(args[2])
			argsMembers := strings.Split((string(args[3])), ",")

			u, err := strconv.ParseUint(argsLimit, 10, 16)

			if len(argsMembers) > int(u) {
				fmt.Println("Number of members is larger than limitation.")
				return err
			}
			//var argsMemstate []types.MemberState
			//
			//
			//for _, s := range argsMembers {
			//	argsMemstate = append(argsMemstate, types.MemberState{
			//		Member: s,
			//		Auth:   false,
			//	})
			//}



			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateJoblist(cliCtx.GetFromAddress(), 
											argsJobcreator, 
											argsTags, 
											argsLimit,
											argsMembers)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}



////////////////////////////////////////////////////////////////////////////////

// func GetCmdCreateJoblistjson(cdc *codec.Codec) *cobra.Command {
// 	json_description := `Json basic format must like this : 
// {
// 	"Jobcreator" : "cosmos......",
// 	"Tags" : ["price", "size"],
// 	"Limit" : 3 ,
// 	"Members" : ["cosmos......", "cosmos......", "cosmos......"]
// }`
// 	return &cobra.Command{
// 		Use:   "create-joblist-json <Path to *.json>",
// 		Short: "Creates a new joblist with json file",
// 		//Long : multiline,
// 		Long: json_description,
// 		Args:  cobra.MinimumNArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			// jsonFile, err := os.Open(args[0])

// 			// fmt.Println(args[0],jsonFile)

// 			// //return err
// 			argsJobcreator := string(args[0])
// 			argsTags := string(args[1])
// 			argsLimit := string(args[2])
// 			argsMembers := string(args[3])
      
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
// 			msg := types.NewMsgCreateJoblist(cliCtx.GetFromAddress(), 
// 											argsJobcreator, 
// 											argsTags, 
// 											argsLimit, 
// 											argsMembers)
// 			err := msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}
// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}
// }