// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     string `gorm:"type:varchar(191)"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

type User struct {
	ID   string `gorm:"type:varchar(191)"`
	UserName     string `json:"UserName"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
}

type InputLogin struct {
	Email    string `json:"email"`
	Passowrd string `json:"passowrd"`
}

type InputRegister struct {
	UserName     string `json:"UserName"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
}

type InputTodo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
