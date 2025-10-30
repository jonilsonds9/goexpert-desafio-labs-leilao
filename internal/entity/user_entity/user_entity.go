package user_entity

import (
	"context"
	"github.com/jonilsonds9/goexpert-desafio-labs-leilao/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(
		ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
