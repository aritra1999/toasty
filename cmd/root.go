package cmd

import (
	"log"
	"toasty/api/ping"

	"github.com/gin-gonic/gin"
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
		initGinServer()
	}
}

func initGinServer() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/ping", ping.Ping)

	r.Run("localhost:8080")
}
