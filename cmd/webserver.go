/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andrerampanelli/hexagonal-arch/adapters/web/server"
	"github.com/spf13/cobra"
)

var webserverCmd = &cobra.Command{
	Use: "webserver",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewWebServer(productService)
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(webserverCmd)
}
