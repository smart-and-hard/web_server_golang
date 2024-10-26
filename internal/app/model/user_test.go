package model_test

import (
	"testing"
	"url-shortener/internal/app/model"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "not_valid_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "5555"
				return u
			},

			isValid: false,
		},
		{
			name: "not_valid_email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "xyu@"
				return u
			},

			isValid: false,
		},
	}

	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			if ts.isValid {
				assert.NoError(t, ts.u().Validate())
			} else {
				assert.Error(t, ts.u().Validate())
			}
		})
	}
}

func TestUser_BeforCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
