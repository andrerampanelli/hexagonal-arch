package cmd

import (
	"database/sql"
	"os"

	appDb "github.com/andrerampanelli/hexagonal-arch/adapters/db"
	"github.com/andrerampanelli/hexagonal-arch/application/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")
var productDbAdapter = appDb.NewProductDb(db)
var productService = service.NewProductService(productDbAdapter)

var rootCmd = &cobra.Command{
	Use: "hexagonal-arch",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
