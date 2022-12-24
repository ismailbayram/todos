package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
)

func (s *ToDoTestSuite) TestNewUserRepository() {
	ur := NewUserRepository(s.db.DBConn)
	assert.Equal(s.T(), ur.db, s.db.DBConn)
}

func (s *ToDoTestSuite) TestCreate() {
	ur := NewUserRepository(s.db.DBConn)

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
	_ = s.db.DBConn.Table("users").Count(&count)
	assert.Equal(s.T(), int64(2), count)
}

func (s *ToDoTestSuite) TestGetByID() {
	ur := NewUserRepository(s.db.DBConn)
	created, err := ur.Create("test_id", "123456", false)
	assert.Nil(s.T(), err)

	user, err := ur.GetByID(created.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), created.ID, user.ID)
}

func (s *ToDoTestSuite) TestGetByUsername() {
	ur := NewUserRepository(s.db.DBConn)
	_, err := ur.Create("test_username", "123456", false)
	assert.Nil(s.T(), err)

	user, err := ur.GetByUsername("test_username")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "test_username", user.Username)
}

func (s *ToDoTestSuite) TestDeactivate() {
	ur := NewUserRepository(s.db.DBConn)
	user, err := ur.Create("deactivated", "123456", true)
	if err != nil {
		panic(err)
	}
	err = ur.Deactivate(user)
	assert.Nil(s.T(), err)
	assert.False(s.T(), user.IsActive)

	result := s.db.DBConn.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.False(s.T(), user.IsActive)
}

func (s *ToDoTestSuite) TestActivate() {
	ur := NewUserRepository(s.db.DBConn)
	user, err := ur.Create("activated", "123456", true)
	if err != nil {
		panic(err)
	}
	err = ur.Activate(user)
	assert.Nil(s.T(), err)
	assert.True(s.T(), user.IsActive)

	result := s.db.DBConn.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.True(s.T(), user.IsActive)
}

func (s *ToDoTestSuite) TestMakeAdmin() {
	ur := NewUserRepository(s.db.DBConn)
	user, err := ur.Create("admin", "123456", false)
	assert.Nil(s.T(), err)

	err = ur.MakeAdmin(user)
	assert.Nil(s.T(), err)
	assert.True(s.T(), user.IsAdmin)

	result := s.db.DBConn.Where("id = ?", user.ID).First(&user)
	assert.Nil(s.T(), result.Error)
	assert.True(s.T(), user.IsActive)
}
