package models

type User struct {
	ID       int    `json:"-"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TodoList struct {
	ID          int `json:"id"`
	UserID      int
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// type TodoItem struct {
// 	Title       string `json:"title" binding:"required"`

// }
