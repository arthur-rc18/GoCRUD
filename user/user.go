package user

type User struct {
	Id       int    `json:"id" `
	Nome     string `json:"nome" `
	Email    string `json:"email" `
	Senha    string `json:"senha" `
	Telefone string `json:"telefone" `
}
