package user

import (
	"git.01.alem.school/Kusbek/social-network/entity"
)

//Reader ...
type Reader interface {
	Find(nickmail string) (*entity.User, error)
	Get(id int) (*entity.User, error)
}

//Writer ...
type Writer interface {
	ChangeVisibility(id, isPublic int) (int, error)
	Create(user *entity.User) (entity.ID, error)
}

//Repository ...
type Repository interface {
	Reader
	Writer
}

//UseCase ...
type UseCase interface {
	CreateUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password string) (*entity.User, error)
	FindUser(nickmail string) (*entity.User, error)
	GetUser(id int) (*entity.User, error)
	ChangeVisibility(id int, isPublic bool) (int, error)
}
