package cmd

import (
	"github.com/spf13/cobra"
)

var cmdComputeTemplate = &cobra.Command{
	Use:     "template",
	Short:   "Compute templates management",
}

func init() {
	cmdCompute.AddCommand(cmdComputeTemplate)
}
