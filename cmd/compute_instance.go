package cmd

import (
	"github.com/spf13/cobra"
)

var cmdComputeInstance = &cobra.Command{
	Use:     "instance",
	Short:   "Compute instances management",
}

func init() {
	cmdCompute.AddCommand(cmdComputeInstance)
}
