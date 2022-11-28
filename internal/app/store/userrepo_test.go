package store_test

import (
	"github.com/JadesHeart/http-rest-api/internal/app/model"
	"github.com/JadesHeart/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, store.NewConfig().DatabaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "user111@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, store.NewConfig().DatabaseURL)
	defer teardown("users")

	email := "user111@example.org"

	_, err := s.User().FindByName(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email: "user111@example.org",
	})
	u, err := s.User().FindByName(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
