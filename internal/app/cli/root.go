package cli

import (
	"github.com/spf13/cobra"
	"log"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd())

	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
