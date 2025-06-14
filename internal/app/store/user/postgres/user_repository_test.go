package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vo1dFl0w/taskmanager-api/internal/app/model"
	"github.com/vo1dFl0w/taskmanager-api/internal/app/store"
	"github.com/vo1dFl0w/taskmanager-api/internal/app/store/user/postgres"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := postgres.TestDB(t, databaseURL)
	defer teardown("users")

	s := postgres.New(db)
	u := model.TestUser(t)

	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := postgres.TestDB(t, databaseURL)
	defer teardown("users")

	s := postgres.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}