package cmd

import (
	"github.com/spf13/cobra"
)

var cmdCompute = &cobra.Command{
	Use:     "compute",
	Short:   "Compute services management",
}

func init() {
	rootCommand.AddCommand(cmdCompute)
}
