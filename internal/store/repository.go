package store

import "url-shortener/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
