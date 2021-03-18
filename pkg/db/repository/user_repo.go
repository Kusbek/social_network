package repository

import (
	"database/sql"

	"git.01.alem.school/Kusbek/social-network/entity"
)

//UserRepository ...
type UserRepository struct {
	db *sql.DB
}

//NewUserRepository ...
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//Create ...
func (r *UserRepository) Create(user *entity.User) (entity.ID, error) {
	res, err := r.db.Exec(
		`
		INSERT INTO users (username, email, first_name, last_name, birth_date, path_to_photo, password)
		VALUES($1,$2,$3,$4,$5,$6,$7)
		`, user.Username, user.Email, user.FirstName, user.LastName, user.BirthDate, user.PathToPhoto, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//Find ...
func (r *UserRepository) Find(nickmail string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.QueryRow(`SELECT id, username, password FROM users WHERE username=$1 or email=$1`, nickmail).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
