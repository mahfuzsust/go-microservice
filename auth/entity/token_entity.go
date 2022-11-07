package entity

type TokenEntity struct {
	Algorithm string `json:alg`
	Type      string `json:typ`
	Token     string `json:token`
}
