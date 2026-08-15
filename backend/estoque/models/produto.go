package models

type Produto struct {
	Codigo    string  `json:"codigo"`
	Descricao string  `json:"descricao"`
	Saldo     float64 `json:"saldo"`
}