package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "wagu"}
	cmd.AddCommand(irCmd)
	cmd.AddCommand(genCmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
