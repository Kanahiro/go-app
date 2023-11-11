package service

import (
	"gin/model"

	"github.com/jmoiron/sqlx"
)

type TodoService struct {
	db sqlx.DB
}

func NewTodoService(db sqlx.DB) *TodoService {
	return &TodoService{db: db}
}

func (s TodoService) GetTodoList() []model.Todo {
	todoList := []model.Todo{}

	rows, err := s.db.Queryx("SELECT * FROM todo")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		todo := model.Todo{}
		err := rows.StructScan(&todo)
		if err != nil {
			panic(err)
		}
		todoList = append(todoList, todo)
	}

	return todoList
}

func (s TodoService) CreateTodo(text string, done bool) model.Todo {
	row, err := s.db.NamedQuery(`INSERT INTO todo (text,done) VALUES (:text,:done) RETURNING id, text, done`,
		map[string]interface{}{
			"text": text,
			"done": done,
		})

	if err != nil {
		panic(err)
	}

	newTodo := model.Todo{}
	row.Next()
	err = row.StructScan(&newTodo)
	if err != nil {
		panic(err)
	}

	return newTodo
}

func (s TodoService) UpdateTodoByID(id int, text string, done bool) model.Todo {
	row, err := s.db.NamedQuery(`UPDATE todo SET text = :text, done = :done WHERE id = :id RETURNING id, text, done`,
		map[string]interface{}{
			"id":   id,
			"text": text,
			"done": done,
		})

	if err != nil {
		panic(err)
	}

	todo := model.Todo{}
	row.Next()
	err = row.StructScan(&todo)
	if err != nil {
		panic(err)
	}

	return todo
}

func (s TodoService) DeleteTodoByID(id int) bool {
	_, err := s.db.NamedQuery(`DELETE FROM todo WHERE id = :id`,
		map[string]interface{}{
			"id": id,
		})

	if err != nil {
		panic(err)
	}

	return true
}
