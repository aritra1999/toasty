package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Cmd() {
	cmd := &cobra.Command{
		Use: "toasty",
		Run: run(),
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}

func run() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		log.Println("Running server")
	}
}
