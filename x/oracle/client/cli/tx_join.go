package cli

import (
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/spf13/cobra"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdJoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join [enclave-report-path] [enc-pub-key-path]",
		Short: "Broadcast message join",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnclaveReportPath := args[0]
			enclaveReport, err := ioutil.ReadFile(argEnclaveReportPath)
			if err != nil {
				return err
			}

			argEncPubKeyPath := args[1]
			encPubKey, err := ioutil.ReadFile(argEncPubKeyPath)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgJoin(
				clientCtx.GetFromAddress().String(),
				enclaveReport,
				&secp256k1.PubKey{Key: encPubKey},
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
