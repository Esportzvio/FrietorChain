package version

import (
	"github.com/esportzvio/frietorchain/command"
	"github.com/esportzvio/frietorchain/versioning"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Returns the current Frietor Saar version",
		Args:  cobra.NoArgs,
		Run:   runCommand,
	}
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	outputter.SetCommandResult(
		&VersionResult{
			Version:   "v0.1.0-beta",
			Commit:    versioning.Commit,
			Branch:    "release/0.1.0-beta",
			BuildTime: "/01/2024",
		},
	)
}
