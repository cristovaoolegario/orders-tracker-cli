// Package cmd /*
package cmd

import (
	"errors"
	"fmt"

	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/cli/correios"
	"github.com/spf13/cobra"
)

// correiosCmd represents the correios command
var correiosCmd = &cobra.Command{
	Use:   "correios",
	Short: "Track an order from correios API",
	Long:  "A longer description for tracking an order from correios API",
	Args:  ValidateArgs,
	Run:   CorreiosRun,
}

func ValidateArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("you need to provide an order number")
	}
	return nil
}

func CorreiosRun(cmd *cobra.Command, args []string) {
	orderNumber := args[0]
	old_ui, _ := cmd.Flags().GetBool("old_ui")

	fmt.Print(old_ui)

	if old_ui {
		correiosCmd := correios.ProvideCorreiosCLI(pkg.CorreiosBaseURL)
		correiosCmd.RetrieveOrder(orderNumber)
	}
	correios.RenderBubbleTeaList(orderNumber)
}

func init() {
	rootCmd.AddCommand(correiosCmd)
	correiosCmd.Flags().BoolP("old_ui", "o", false, "Render the old ui")
}
