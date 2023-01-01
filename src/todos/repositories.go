package todos

import (
	"github.com/ismailbayram/todos/src/users"
	"gorm.io/gorm"
)

type ToDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *ToDoRepository {
	return &ToDoRepository{
		db: db,
	}
}

func (r *ToDoRepository) Create(name string, user users.User) (*ToDo, error) {
	toDo := &ToDo{
		Name: name,
		User: user,
	}
	result := r.db.Create(toDo)
	if result.Error != nil {
		return nil, result.Error
	}
	return toDo, nil
}

func (r *ToDoRepository) MakeDone(toDo *ToDo) error {
	toDo.IsDone = true
	return r.db.Save(toDo).Error
}
