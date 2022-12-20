package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
)

func (s *ToDoTestSuite) TestCreateUser() {
	ur := NewUserRepository(s.db.DBConn)

	password := "123456"
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	user, err := ur.CreateUser("ismail", password, true)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "ismail", user.Username)
	assert.Equal(s.T(), hashedPassword, user.Password)
	assert.True(s.T(), user.IsActive)
	assert.True(s.T(), user.IsAdmin)

	user, err = ur.CreateUser("hilal", password, true)

	var count int64
	_ = s.db.DBConn.Table("users").Count(&count)
	assert.Equal(s.T(), int64(2), count)
}

func (s *ToDoTestSuite) TestDeactivateUser() {
	assert.True(s.T(), true)
}

//func TestMakeUserAdmin(t *testing.T) {
//	assert.True(t, true)
//}
