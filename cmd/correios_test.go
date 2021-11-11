package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCorreiosCmd_ValidateArgs(t *testing.T) {
	t.Run("Should return error when there's no args", func(t *testing.T) {
		result := ValidateArgs(&cobra.Command{}, []string{})

		expected := "you need to provide an order number"
		if result == nil && result.Error() != expected {
			t.Fatalf("Should return the error '%s', got '%s'", expected, result)
		}
	})

	t.Run("Should not return error when there's a arg", func(t *testing.T) {
		result := ValidateArgs(&cobra.Command{}, []string{"param"})

		if result != nil {
			t.Fatalf("Shouldn't return errors, got '%s'", result)
		}
	})
}

// func TestCorreiosCmd_CorreiosRun(t *testing.T) {
// 	t.Run("Should render OLD UI when there's the flag -o", func(t *testing.T) {
// 		cmd := cobra.Command{}
// 		cmd.Flags().BoolP("old_ui", "o", true, "")

// 		//CorreiosRun(&cmd, []string{"test"})

// 	})
// }
