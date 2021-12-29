package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdComputeInstanceGet = &cobra.Command{
	Use:   "get",
	Short: "Get/show compute instance",
	Args:  cobra.ExactArgs(1),
	RunE:  ComputeInstanceGet,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceGet, FLAG_ZONE)
	cmdComputeInstance.AddCommand(cmdComputeInstanceGet)
}

func ComputeInstanceGet(cmd *cobra.Command, args []string) error {
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
	resp, err := client.GetInstanceWithResponse(apiContext, id)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceGet] FATAL: Failed to get instance; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
