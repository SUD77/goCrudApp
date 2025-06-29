package model

type User struct{

	ID int `json:"id"`
	Name string `json:"name`
	Email string `json:"email"`
}

const (
	InserUser = `INSERT INTO users(name,email) VALUES($1,$2) RETURNING id`
	GetUser = `SELECT id,name,email FROM users WHERE id=$1`
	ListUsers = `SELECT id,name,email FROM users`
	UpdateUser = `UPDATE users SET name=$1,email=$2 WHERE id=$3`
	DeleteUser = `DELETE FROM users WHERE id=$1`
)