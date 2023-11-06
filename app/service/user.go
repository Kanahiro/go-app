package service

import (
	"gin/graph/model"
	"gin/infra"
)

func GetUsers() ([]*model.User, error) {
	db := infra.Db()
	rows, _ := db.Queryx("SELECT * FROM myuser")

	var users []*model.User
	// db-user to model-user
	for rows.Next() {
		var user model.User
		rows.StructScan(&user)
		users = append(users, &user)
	}

	return users, nil
}

func CreateUser(name string, email *string) (*model.User, error) {
	db := infra.Db()
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO myuser (name, email) VALUES ($1, $2)", name, email)
	tx.Commit()

	return &model.User{
		Name:  name,
		Email: email,
	}, nil
}
