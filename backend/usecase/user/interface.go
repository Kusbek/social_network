package user

import (
	"git.01.alem.school/Kusbek/social-network/backend/entity"
)

type Reader interface {
	Find(nickmail string) (*entity.User, error)
}

type Writer interface {
	Create(user *entity.User) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password string) (entity.ID, error)
	FindUser(nickmail, password string) (*entity.User, error)
}
