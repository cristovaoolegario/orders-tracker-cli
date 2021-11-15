package dto

// CorreiosResponse represents the entire object of correios API response
type CorreiosResponse struct {
	Objects []Object `json:"objeto,omitempty"`
	Search  string   `json:"pesquisa"`
	Amount  string   `json:"quantidade"`
	Result  string   `json:"resultado"`
	Version string   `json:"versao"`
}

// Event represents an event of correios API response
type Event struct {
	DestinationCode string     `json:"cepDestino"`
	CreationDate    string     `json:"criacao"`
	Data            string     `json:"data"`
	PostDate        string     `json:"dataPostagem"`
	Description     string     `json:"descricao"`
	Time            string     `json:"hora"`
	Status          string     `json:"status"`
	Type            string     `json:"tipo"`
	Unit            Unit       `json:"unidade"`
	Details         string     `json:"detalhe,omitempty"`
	Post            PostalType `json:"postagem,omitempty"`
}

// Object represents an object of correios API response
type Object struct {
	Category string  `json:"categoria"`
	Events   []Event `json:"evento"`
	Name     string  `json:"nome"`
	Number   string  `json:"numero"`
	Initials string  `json:"sigla"`
}

// Address represents an address of correios API response
type Address struct {
	District string `json:"bairro"`
	Cep      string `json:"cep"`
	Code     string `json:"codigo"`
	City     string `json:"localidade"`
	Street   string `json:"logradouro"`
	Number   string `json:"numero"`
	Uf       string `json:"uf"`
}

// Unit represents a unit of correios API response
type Unit struct {
	City     string  `json:"cidade"`
	Code     string  `json:"codigo"`
	Address  Address `json:"endereco"`
	Local    string  `json:"local"`
	Sto      string  `json:"sto"`
	UnitType string  `json:"tipounidade"`
	Uf       string  `json:"uf"`
}

// Recipient represents a recipient of correios API response
type Recipient struct {
	Cep string `json:"cep"`
}

// PostalType represents a postal type of correios API response
type PostalType struct {
	Ar              string `json:"ar"`
	Cepdestino      string `json:"cepdestino"`
	Datapostagem    string `json:"datapostagem"`
	Dataprogramada  string `json:"dataprogramada"`
	Dh              string `json:"dh"`
	Mp              string `json:"mp"`
	Peso            string `json:"peso"`
	Prazotratamento string `json:"prazotratamento"`
	Volume          string `json:"volume"`
}
