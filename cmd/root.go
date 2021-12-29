package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/exoscale/openego/openego"
)

var rootCommand = &cobra.Command{
	Use:           "openego",
	Short:         "Exoscale OpenAPI Go Wrapper",
	Version:       openego.VERSION+" ("+openego.COMMIT+")",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {

	// Execution context and OS signals
	rootContext, rootContext_fnCancel := UseRootContext()

	// (signals)
	chSignal := make(chan os.Signal, 1)
	signal.Notify(
		chSignal,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	// (graceful exit)
	defer func() {
		signal.Stop(chSignal)
		rootContext_fnCancel()
	}()

	// (watch signals/termination)
	go func() {
		select {

		case s := <-chSignal:
			// OS signal
			rootContext_fnCancel()
			log.Printf("[cmd.Execute] INFO: OS signal; %s", s)

		case <-rootContext.Done():
			// Context timeout
			log.Print("[cmd.Execute] INFO: Context timeout")

		}
	}()

	// Root command
	if err := rootCommand.Execute(); err != nil {
		log.Fatalf("[cmd.Execute] FATAL: Failed to execute command; %s", err)
	}
}

func JsonToString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("[cmd.JsonToString] ERROR: Failed to parse JSON response; %s", err)
		return ""
	}
	return string(bytes)
}
