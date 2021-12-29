package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	exoOpenApi "github.com/exoscale/egoscale/v2/oapi"
)

var cmdComputeTemplateList = &cobra.Command{
	Use:   "list",
	Short: "List compute templates",
	Args:  cobra.NoArgs,
	RunE:  ComputeTemplateList,
}

func init() {
	defaults := map[string]string{
		"visibility": "private",
	}
	AddCanonicalFlags(cmdComputeTemplateList, FLAG_ZONE)
	AddOpenApiFlags(
		cmdComputeTemplateList,
		&exoOpenApi.ListTemplatesParams{},
		&defaults,
	)
	cmdComputeTemplate.AddCommand(cmdComputeTemplateList)
}

func ComputeTemplateList(cmd *cobra.Command, args []string) error {
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
	params := &exoOpenApi.ListTemplatesParams{}
	ParseOpenApiFlags(cmd, params)

	// Call
	resp, err := client.ListTemplatesWithResponse(apiContext, params)
	if err != nil {
		log.Fatalf("[cmd.runComputeTemplateList] FATAL: Failed to list templates; %s", err)
	}
	fmt.Print(JsonToString(resp.JSON200))

	// Done
	return nil
}
