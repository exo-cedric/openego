package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	exoOpenApi "github.com/exoscale/egoscale/v2/oapi"
)

var cmdComputeInstanceList = &cobra.Command{
	Use:   "list",
	Short: "List compute instances",
	Args:  cobra.NoArgs,
	RunE:  ComputeInstanceList,
}

func init() {
	AddCanonicalFlags(cmdComputeInstanceList, FLAG_ZONE)
	AddOpenApiFlags(
		cmdComputeInstanceList,
		&exoOpenApi.ListInstancesParams{},
		nil,
	)
	cmdComputeInstance.AddCommand(cmdComputeInstanceList)
}

func ComputeInstanceList(cmd *cobra.Command, args []string) error {
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

	// Params
	params := &exoOpenApi.ListInstancesParams{}
	ParseOpenApiFlags(cmd, params)

	// Call
	resp, err := client.ListInstancesWithResponse(apiContext, params)
	if err != nil {
		log.Fatalf("[cmd.runComputeInstanceList] FATAL: Failed to list instances; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
