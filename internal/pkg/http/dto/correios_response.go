package dto

// CorreiosResponse represents the entire object of correios API response
type CorreiosResponse struct {
	Objects []Object `json:"objetos,omitempty"`
	Amount  int      `json:"quantidade,omitempty"`
	Result  string   `json:"resultado,omitempty"`
	Version string   `json:"versao,omitempty"`
}

// Event represents an event of correios API response
type Event struct {
	Code            string    `json:"codigo"`
	Description     string    `json:"descricao"`
	DateTimeCreated string    `json:"dtHrCriado"`
	Objects         []Object  `json:"objetos"`
	Amount          int       `json:"quantidade"`
	Result          string    `json:"resultado"`
	Version         string    `json:"versao"`
	Type            string    `json:"tipo"`
	Unit            Unit      `json:"unidade"`
	IconURL         string    `json:"urlIcone"`
	DestinationUnit Unit      `json:"unidadeDestino,omitempty"`
	Detail          string    `json:"detalhe,omitempty"`
	Recipient       Recipient `json:"destinatario,omitempty"`
}

// Object represents an object of correios API response
type Object struct {
	ObjectCode   string     `json:"codObjeto,omitempty"`
	Message      string     `json:"mensagem,omitempty"`
	ExpectedDate string     `json:"dtPrevista,omitempty"`
	Events       []Event    `json:"eventos,omitempty"`
	Modality     string     `json:"modalidade,omitempty"`
	PostalType   PostalType `json:"tipoPostal,omitempty"`
}

// Address represents an address of correios API response
type Address struct {
	City     string `json:"cidade"`
	State    string `json:"uf"`
	District string `json:"bairro,omitempty"`
	Cep      string `json:"cep,omitempty"`
	Street   string `json:"logradouro,omitempty"`
	Number   string `json:"numero,omitempty"`
}

// Unit represents a unit of correios API response
type Unit struct {
	Address Address `json:"endereco"`
	Type    string  `json:"tipo"`
}

// Recipient represents a recipient of correios API response
type Recipient struct {
	Cep string `json:"cep"`
}

// PostalType represents a postal type of correios API response
type PostalType struct {
	Category    string `json:"categoria"`
	Description string `json:"descricao"`
	Initials    string `json:"sigla"`
}

type TokenType struct {
	Token string `json:"token"`
}
