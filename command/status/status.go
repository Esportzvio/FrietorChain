package status

import (
	"context"

	"github.com/esportzvio/frietorchain/command"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"

	"github.com/esportzvio/frietorchain/command/helper"
	"github.com/esportzvio/frietorchain/server/proto"
)

func GetCommand() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Returns the status of the Frietor Saar client",
		Args:  cobra.NoArgs,
		Run:   runCommand,
	}

	helper.RegisterGRPCAddressFlag(statusCmd)

	return statusCmd
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	statusResponse, err := getSystemStatus(helper.GetGRPCAddress(cmd))
	if err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(&StatusResult{
		ChainID:            statusResponse.Network,
		CurrentBlockNumber: statusResponse.Current.Number,
		CurrentBlockHash:   statusResponse.Current.Hash,
		LibP2PAddress:      statusResponse.P2PAddr,
	})
}

func getSystemStatus(grpcAddress string) (*proto.ServerStatus, error) {
	client, err := helper.GetSystemClientConnection(
		grpcAddress,
	)
	if err != nil {
		return nil, err
	}

	return client.GetStatus(context.Background(), &empty.Empty{})
}
