package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdComputeInstanceStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop compute instance",
	Args:  cobra.ExactArgs(1),
	RunE:  ComputeInstanceStop,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceStop, FLAG_ZONE)
	cmdComputeInstance.AddCommand(cmdComputeInstanceStop)
}

func ComputeInstanceStop(cmd *cobra.Command, args []string) error {
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
	resp, err := client.StopInstanceWithResponse(apiContext, id)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceStop] FATAL: Failed to stop instance; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
