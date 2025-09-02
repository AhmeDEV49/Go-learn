package repository

import "github.com/ahmedev49/go-clean-architecture/internal/entity"

// UserRepository defines the contract for persisting users.
// Any implementation (Postgres, MySQL, in-memory) must satisfy this interface.
type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	GetAll() ([]entity.User, error)
}
