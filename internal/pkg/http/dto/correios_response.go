package dto

type CorreiosResponse struct {
	Objetos    []Objeto `json:"objetos,omitempty"`
	Quantidade int      `json:"quantidade,omitempty"`
	Resultado  string   `json:"resultado,omitempty"`
	Versao     string   `json:"versao,omitempty"`
}

type Evento struct {
	Codigo         string       `json:"codigo"`
	Descricao      string       `json:"descricao"`
	DtHrCriado     string       `json:"dtHrCriado"`
	Objetos        []Objeto     `json:"objetos"`
	Quantidade     int          `json:"quantidade"`
	Resultado      string       `json:"resultado"`
	Versao         string       `json:"versao"`
	Tipo           string       `json:"tipo"`
	Unidade        Unidade      `json:"unidade"`
	UrlIcone       string       `json:"urlIcone"`
	UnidadeDestino Unidade      `json:"unidadeDestino,omitempty"`
	Detalhe        string       `json:"detalhe,omitempty"`
	Destinatario   Destinatario `json:"destinatario,omitempty"`
}

type Objeto struct {
	CodObjeto                  string     `json:"codObjeto,omitempty"`
	Mensagem                   string     `json:"mensagem,omitempty"`
	DtPrevista                 string     `json:"dtPrevista,omitempty"`
	Eventos                    []Evento   `json:"eventos,omitempty"`
	Modalidade                 string     `json:"modalidade,omitempty"`
	TipoPostal                 TipoPostal `json:"tipoPostal,omitempty"`
	HabilitaAutoDeclaracao     bool       `json:"habilitaAutoDeclaracao,omitempty"`
	PermiteEncargoImportacao   bool       `json:"permiteEncargoImportacao,omitempty"`
	HabilitaPercorridaCarteiro bool       `json:"habilitaPercorridaCarteiro,omitempty"`
	BloqueioObjeto             bool       `json:"bloqueioObjeto,omitempty"`
	PossuiLocker               bool       `json:"possuiLocker,omitempty"`
	HabilitaLocker             bool       `json:"habilitaLocker,omitempty"`
	HabilitaCrowdshipping      bool       `json:"habilitaCrowdshipping,omitempty"`
}

type Endereco struct {
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Bairro     string `json:"bairro,omitempty"`
	Cep        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Numero     string `json:"numero,omitempty"`
}

type Unidade struct {
	Endereco Endereco `json:"endereco"`
	Tipo     string   `json:"tipo"`
}

type Destinatario struct {
	Cep string `json:"cep"`
}
type TipoPostal struct {
	Categoria string `json:"categoria"`
	Descricao string `json:"descricao"`
	Sigla     string `json:"sigla"`
}
