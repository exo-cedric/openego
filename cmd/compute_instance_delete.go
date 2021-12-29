package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdComputeInstanceDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete compute instance",
	Args:  cobra.ExactArgs(1),
	RunE:  ComputeInstanceDelete,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceDelete, FLAG_ZONE)
	cmdComputeInstance.AddCommand(cmdComputeInstanceDelete)
}

func ComputeInstanceDelete(cmd *cobra.Command, args []string) error {
	// Client
	client, err := NewClient()
	if err != nil {
		return err
	}

	// Arguments
	id := args[0]

	// Context
	zone, err := cmd.Flags().GetString("zone")
	if err != nil {
		return err
	}
	apiContext := NewApiContext(zone)

	// Call
	resp, err := client.DeleteInstanceWithResponse(apiContext, id)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceDelete] FATAL: Failed to delete instance; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
