package store_test

import (
	"testing"

	"github.com/ibadullaev-inc4/Tutorial_REST_Middleware/internal/model"
	"github.com/ibadullaev-inc4/Tutorial_REST_Middleware/store"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@example.com",
		EncryptedPassword: "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email:             "user@example.org",
		EncryptedPassword: "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
