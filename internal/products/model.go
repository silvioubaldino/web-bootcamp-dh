package products

type User struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float64 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"dataCriacao"`
}

type Request struct {
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float64 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"dataCriacao"`
}

func ReqToUser(r Request) User {
	return User{
		Nome:        r.Nome,
		Sobrenome:   r.Sobrenome,
		Email:       r.Email,
		Idade:       r.Idade,
		Altura:      r.Altura,
		Ativo:       r.Ativo,
		DataCriacao: r.DataCriacao,
	}
}
