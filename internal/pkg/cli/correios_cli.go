package cli

import (
	"fmt"
	"orders-tracker-cli/internal/pkg/http/services"
)

type CorreiosCLI struct {
	service services.ICorreiosService
}

func ProvideCorreiosCLI(baseUrl string) CorreiosCLI {
	return CorreiosCLI{service: services.ProvideCorreiosService(baseUrl)}
}

func (cli *CorreiosCLI) RetrieveOrder(orderNumber string) {
	response, err := cli.service.FindOrderByNumber(orderNumber)
	if err == nil {
		for _, evento := range response.Objetos[0].Eventos {
			fmt.Println("🚚 - " + evento.Descricao)
			fmt.Println("⏱ - " + evento.DtHrCriado)
			fmt.Println()
		}
	} else {
		fmt.Printf("❌ - %s\n", err.Error())
	}
}
