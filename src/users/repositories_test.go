package users

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ismailbayram/todos/src/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	tests.AppTestSuite
}

func TestUserTestSuite(t *testing.T) {
	userTestSuite := new(UserTestSuite)
	userTestSuite.Models = []interface{}{&User{}}
	suite.Run(t, userTestSuite)
}

func (s *UserTestSuite) TestNewUserRepository() {
	ur := NewUserRepository(s.DB)
	assert.Equal(s.T(), ur.db, s.DB)
}

func (s *UserTestSuite) TestCreate() {
	ur := NewUserRepository(s.DB)

	password := "123456"
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	user, err := ur.Create("ismail", password, true)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "ismail", user.Username)
	assert.Equal(s.T(), hashedPassword, user.Password)
	assert.True(s.T(), user.IsActive)
	assert.True(s.T(), user.IsAdmin)

	user, err = ur.Create("hilal", password, true)

	var count int64
	_ = s.DB.Table("users").Count(&count)
	assert.Equal(s.T(), int64(2), count)
}

func (s *UserTestSuite) TestGetByID() {
	ur := NewUserRepository(s.DB)
	created, err := ur.Create("test_id", "123456", false)
	assert.Nil(s.T(), err)

	user, err := ur.GetByID(created.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), created.ID, user.ID)

	user2, err2 := ur.GetByID(12121231132123)
	assert.NotNil(s.T(), err2)
	assert.Equal(s.T(), uint(0), user2.ID)
}

func (s *UserTestSuite) TestGetByUsername() {
	ur := NewUserRepository(s.DB)
	_, err := ur.Create("test_username", "123456", false)
	assert.Nil(s.T(), err)

	user, err := ur.GetByUsername("test_username")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "test_username", user.Username)
}

func (s *UserTestSuite) TestDeactivate() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("deactivated", "123456", true)
	if err != nil {
		panic(err)
	}
	err = ur.Deactivate(user)
	assert.Nil(s.T(), err)
	assert.False(s.T(), user.IsActive)

	result := s.DB.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.False(s.T(), user.IsActive)
}

func (s *UserTestSuite) TestActivate() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("activated", "123456", true)
	if err != nil {
		panic(err)
	}
	err = ur.Activate(user)
	assert.Nil(s.T(), err)
	assert.True(s.T(), user.IsActive)

	result := s.DB.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.True(s.T(), user.IsActive)
}

func (s *UserTestSuite) TestMakeAdmin() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("admin", "123456", false)
	assert.Nil(s.T(), err)

	err = ur.MakeAdmin(user)
	assert.Nil(s.T(), err)
	assert.True(s.T(), user.IsAdmin)

	result := s.DB.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.True(s.T(), user.IsActive)
}

func (s *UserTestSuite) TestCreateToken() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("token", "123456", true)
	assert.Nil(s.T(), err)
	token, err := ur.CreateToken(user, s.Config.SecretKey)
	assert.Nil(s.T(), err)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", "token", s.Config.SecretKey)))
	expectedToken := hex.EncodeToString(h.Sum(nil))
	assert.Equal(s.T(), expectedToken, token)
}

func (s *UserTestSuite) TestGetByToken() {
	ur := NewUserRepository(s.DB)
	user, err := ur.Create("token", "123456", true)
	assert.Nil(s.T(), err)
	token, err := ur.CreateToken(user, s.Config.SecretKey)
	assert.Nil(s.T(), err)

	expectedUser, err := ur.GetByToken(token)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedUser.ID, user.ID)

	expectedUser, err = ur.GetByToken("asdasdasdasd")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), uint(0), expectedUser.ID)
}
