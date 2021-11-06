// Package cmd /*
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"orders-tracker-cli/internal/pkg"
	"orders-tracker-cli/internal/pkg/cli"
)

// correiosCmd represents the correios command
var correiosCmd = &cobra.Command{
	Use:   "correios",
	Short: "Track a order from correios API.",
	Long:  "Track a order from correios API.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0{
			return errors.New("You need to provide an order number!")
		}
		correiosCmd := cli.ProvideCorreiosCLI(pkg.CorreiosBaseUrl)
		correiosCmd.RetrieveOrder(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(correiosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// correiosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// correiosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
