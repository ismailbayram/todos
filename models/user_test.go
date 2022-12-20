package models

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *ToDoTestSuite) TestCreateUser() {
	password := "123456"
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := string(h.Sum(nil))
	user, err := CreateUser(s.db, "ismail", password, true)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "ismail", user.Username)
	assert.Equal(s.T(), hashedPassword, user.Password)
	assert.True(s.T(), user.IsActive)
	assert.True(s.T(), user.IsAdmin)
}

func (s *ToDoTestSuite) TestDeactivateUser() {
	assert.True(s.T(), true)
}

func TestMakeUserAdmin(t *testing.T) {
	assert.True(t, true)
}
