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

func (r *ToDoRepository) GetByUserID(userID uint) ([]ToDo, error) {
	toDos := []ToDo{}
	result := r.db.Order("id asc").Where("user_id = ?", userID).Find(&toDos)
	return toDos, result.Error
}

func (r *ToDoRepository) GetByID(id int) (*ToDo, error) {
	var toDo ToDo
	result := r.db.Where("id = ?", id).First(&toDo)
	return &toDo, result.Error
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

func (r *ToDoRepository) Update(toDo *ToDo) error {
	return r.db.Save(toDo).Error
}
