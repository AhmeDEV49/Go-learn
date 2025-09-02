package infrastructure

import (
	"database/sql"
	"github.com/ahmedev49/go-clean-architecture/internal/entity"
	"github.com/ahmedev49/go-clean-architecture/internal/repository"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) repository.UserRepository {
	return &UserRepositoryPostgres{db: db}
}

func (r UserRepositoryPostgres) Create(user *entity.User) error {
	err := r.db.QueryRow(
		"INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r UserRepositoryPostgres) GetByID(id int64) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepositoryPostgres) GetAll() ([]entity.User, error) {
	var users []entity.User

	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
