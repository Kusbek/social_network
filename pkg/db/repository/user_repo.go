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
		INSERT INTO users (username, email, first_name, last_name, birth_date, path_to_photo, about_me, password)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)
		`,
		user.Username,
		user.Email,
		user.FirstName,
		user.LastName,
		entity.TimeToString(user.BirthDate),
		user.PathToPhoto,
		user.AboutMe,
		user.Password)
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

//Get ...
func (r *UserRepository) Get(id int) (*entity.User, error) {
	user := &entity.User{}
	var birthDate string
	err := r.db.QueryRow(`SELECT id, username, email, first_name, last_name, about_me, path_to_photo, birth_date, is_public FROM users WHERE id=$1`, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AboutMe,
		&user.PathToPhoto,
		&birthDate,
		&user.IsPublic,
	)
	user.BirthDate = entity.StringToTime(birthDate)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) ChangeVisibility(userID, isPublic int) (int, error) {

	res, err := r.db.Exec(`UPDATE users SET is_public=$1 WHERE id=$2`, isPublic, userID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
