package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/youngjoon-lee/dhub/x/oracle/types"
)

func CmdListOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-oracles",
		Short: "list all oracles",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOracleRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OracleAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-oracle [operator-address]",
		Short: "shows an oracle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argOperatorAddress := args[0]

			params := &types.QueryGetOracleRequest{
				OperatorAddress: argOperatorAddress,
			}

			res, err := queryClient.Oracle(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOraclePubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-oracle-pubkey",
		Short: "shows an oracle pubkey",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.OraclePubKey(context.Background(), &types.QueryGetOraclePubKeyRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
