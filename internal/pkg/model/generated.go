// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// Мутации связанные с авторизацией
type AuthorizationMutations struct {
	//  Регистрация
	Signup *SignupOutput `json:"signup"`
}

type Mutation struct {
}

type Query struct {
}

// Входные параметры для регистрации
type SignupInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Выходные параметры для регистрации
type SignupOutput struct {
	Result int `json:"result"`
}

type TestInput struct {
	Test string `json:"test"`
}

type TestOutput struct {
	Test string `json:"test"`
}
