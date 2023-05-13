package models

type CreateUserDTO struct {
	Login          string `json:"login" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"password_repeat" binding:"required"`
}

type SignInUserDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `json:"-"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TodoList struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TodoElement struct {
	ID         int    `json:"id"`
	TodoListID int    `json:"todo_list_id"`
	Title      string `json:"title" binding:"required"`
	Status     string `json:"status_id"`
}
