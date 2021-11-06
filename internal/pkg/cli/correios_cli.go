package cli

import (
	"fmt"
	"github.com/cristovaoolegario/orders-tracker-cli/internal/pkg/http/services"
)

// CorreiosCLI provides the cli validation for correios command
type CorreiosCLI struct {
	service services.ICorreiosService
}

// ProvideCorreiosCLI provides a CorreiosCLI
func ProvideCorreiosCLI(baseURL string) CorreiosCLI {
	return CorreiosCLI{service: services.ProvideCorreiosService(baseURL)}
}

// RetrieveOrder prints the order data on the terminal
func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	if err == nil {
		for _, event := range response.Objects[0].Events {
			fmt.Println("🚚 - " + event.Description)
			fmt.Println("⏱ - " + event.DateTimeCreated)
			fmt.Println()
		}
	} else {
		fmt.Printf("❌ - %s\n", err.Error())
	}
}
