// Package cmd /*
package cmd

import (
	"errors"

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
	RunE:  CorreiosRunE,
}

func ValidateArgs(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("you need to provide an order number")
	}
	return nil
}

func CorreiosRunE(cmd *cobra.Command, args []string) error {
	orderNumber := args[0]
	old_ui, _ := cmd.Flags().GetBool("old_ui")

	if old_ui {
		correiosCmd := correios.ProvideCorreiosCLI(pkg.CorreiosBaseURL, pkg.CorreiosValidationUrl)
		correiosCmd.RetrieveOrder(orderNumber)
		return nil
	}
	correios.RenderBubbleTeaList(orderNumber)
	return nil
}

func init() {
	rootCmd.AddCommand(correiosCmd)
	correiosCmd.Flags().BoolP("old_ui", "o", false, "Render the old ui")
}
