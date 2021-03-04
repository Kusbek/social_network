package user

import (
	"git.01.alem.school/Kusbek/social-network/backend/entity"
)

//Reader ...
type Reader interface {
	Find(nickmail string) (*entity.User, error)
}

//Writer ...
type Writer interface {
	Create(user *entity.User) (entity.ID, error)
}

//Repository ...
type Repository interface {
	Reader
	Writer
}

//UseCase ...
type UseCase interface {
	CreateUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password string) (entity.ID, error)
	FindUser(nickmail string) (*entity.User, error)
}
