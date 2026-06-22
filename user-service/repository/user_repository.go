package repository

import "gorm.io/gorm"

// User: Get By email, get by id, Create, Update, Delete, Get by rolename, Get all
// Role: Get by id, Create, Update, Delete, Get all
// UserRole: Assign User to role, getUserRoleById, get all user role, edit assign user to role

type UserRepositoryInterface interface {
	GetByEmail(email string) (*User, error)
	GetById(id int) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id int) error
	GetByRolename(rolename string) (*User, error)
	GetAll() ([]User, error)
}

type UserRepository struct {
	db *gorm.DB
}
