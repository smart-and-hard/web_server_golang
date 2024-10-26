package teststore

import (
	"url-shortener/internal/app/model"
	"url-shortener/internal/store"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.UserRepository
}
