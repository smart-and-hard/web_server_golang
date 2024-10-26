package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "kotelnikov@test.ru",
		Password: "123456",
	}
}
