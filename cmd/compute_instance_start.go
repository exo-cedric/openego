package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	exoOpenApi "github.com/exoscale/egoscale/v2/oapi"
)

var cmdComputeInstanceStart = &cobra.Command{
	Use:   "start",
	Short: "Start compute instance",
	Args:  cobra.ExactArgs(1),
	RunE:  ComputeInstanceStart,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceStart, FLAG_ZONE)
	AddOpenApiFlags(
		cmdComputeInstanceStart,
		&exoOpenApi.StartInstanceJSONRequestBody{},
		nil,
	)
	cmdComputeInstance.AddCommand(cmdComputeInstanceStart)
}

func ComputeInstanceStart(cmd *cobra.Command, args []string) error {
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

	// Params
	params := &exoOpenApi.StartInstanceJSONRequestBody{}
	ParseOpenApiFlags(cmd, params)

	// Call
	resp, err := client.StartInstanceWithResponse(apiContext, id, *params)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceStart] FATAL: Failed to start instance; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
