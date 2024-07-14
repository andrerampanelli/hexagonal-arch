package cmd

import (
	"github.com/andrerampanelli/hexagonal-arch/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var id string
var name string
var price float64

var cliCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(productService, action, id, name, price)
		if err != nil {
			cmd.PrintErr(err)
		}
		cmd.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "", "Action to be executed")
	cliCmd.Flags().StringVarP(&id, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&name, "name", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&price, "price", "p", 0.0, "Product Price")
}
