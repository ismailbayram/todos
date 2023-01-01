package todos

import (
	"github.com/ismailbayram/todos/src/users"
	"github.com/stretchr/testify/assert"
)

func (s *ToDoTestSuite) TestNewUserRepository() {
	tdr := NewToDoRepository(s.DB)
	assert.Equal(s.T(), tdr.db, s.DB)
}

func (s *ToDoTestSuite) TestCreate() {
	tdr := NewToDoRepository(s.DB)
	ur := users.NewUserRepository(s.DB)

	user, _ := ur.Create("ismail", "123456", true)

	toDo, err := tdr.Create("First To Do", *user)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "First To Do", toDo.Name)
	assert.Equal(s.T(), *user, toDo.User)
	assert.False(s.T(), toDo.IsDone)

	var count int64
	_ = s.DB.Table("to_dos").Count(&count)
	assert.Equal(s.T(), int64(1), count)
}

func (s *ToDoTestSuite) TestMakeDone() {
	tdr := NewToDoRepository(s.DB)
	ur := users.NewUserRepository(s.DB)

	user, _ := ur.Create("ismail", "123456", true)

	toDo, err := tdr.Create("First To Do", *user)
	assert.Nil(s.T(), err)
	assert.False(s.T(), toDo.IsDone)

	err = tdr.MakeDone(toDo)
	assert.Nil(s.T(), err)
	assert.True(s.T(), toDo.IsDone)
}
