package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/esportzvio/frietorchain/command/helper"
	"github.com/esportzvio/frietorchain/command/peers"
	"github.com/esportzvio/frietorchain/command/polybftsecrets"
	"github.com/esportzvio/frietorchain/command/server"
	"github.com/esportzvio/frietorchain/command/status"
	"github.com/esportzvio/frietorchain/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Frietor Saar is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		status.GetCommand(),
		polybftsecrets.GetCommand(),
		peers.GetCommand(),
		server.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
