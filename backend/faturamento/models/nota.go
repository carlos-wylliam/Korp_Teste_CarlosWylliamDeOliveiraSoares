package models

type ItemNota struct {
	ProdutoCodigo string  `json:"produtoCodigo"`
	Quantidade    float64 `json:"quantidade"`
}

type Nota struct {
	Numero int        `json:"numero"`
	Status string     `json:"status"`
	Itens  []ItemNota `json:"itens"`
}