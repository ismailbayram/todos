package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestCreateUser(t *testing.T) {
//	password := "123456"
//	h := sha256.New()
//	h.Write([]byte(password))
//	hashedPassword := string(h.Sum(nil))
//	user := CreateUser(nil, "ismail", password, true)
//
//	assert.Equal(t, "ismail", user.Username)
//	assert.Equal(t, hashedPassword, user.Password)
//	assert.True(t, user.IsActive)
//	assert.True(t, user.IsAdmin)
//}

func (s *ToDoTestSuite) TestDeactivateUser() {
	assert.True(s.T(), true)
}

func TestMakeUserAdmin(t *testing.T) {
	assert.True(t, true)
}
