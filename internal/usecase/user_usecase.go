package usecase

import (
	"github.com/ahmedev49/go-clean-architecture/internal/entity"
	"github.com/ahmedev49/go-clean-architecture/internal/repository"
)

// UserUsecase contains the business logic for users.
type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// CreateUser handles the business logic of creating a user.
func (uc *UserUsecase) CreateUser(user *entity.User) error {
	if user.Email == "" {
		return ErrInvalidEmail
	}
	return uc.repo.Create(user)
}

func (uc *UserUsecase) GetUserById(id int64) (*entity.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUsecase) ListUsers() ([]entity.User, error) {
	return uc.repo.GetAll()
}
