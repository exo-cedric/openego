package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdComputeInstanceTypes = &cobra.Command{
	Use:   "types",
	Short: "List compute instance types",
	Args:  cobra.NoArgs,
	RunE:  ComputeInstanceTypes,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceTypes, FLAG_ZONE)
	cmdComputeInstance.AddCommand(cmdComputeInstanceTypes)
}

func ComputeInstanceTypes(cmd *cobra.Command, args []string) error {
	// Client
	client, err := NewClient()
	if err != nil {
		return err
	}

	// Context
	zone, err := cmd.Flags().GetString("zone")
	if err != nil {
		return err
	}
	apiContext := NewApiContext(zone)

	// Call
	resp, err := client.ListInstanceTypesWithResponse(apiContext)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceTypes] FATAL: Failed to list instance types; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
