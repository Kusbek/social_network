package session

import "git.01.alem.school/Kusbek/social-network/backend/entity"

//Reader ...
type Reader interface {
	Get(uuid string) (*entity.User, error)
}

//Writer ...
type Writer interface {
	Create(user *entity.User) string
	Delete(uuid string)
}

//Repository ...
type Repository interface {
	Reader
	Writer
}

//UseCase ...
type UseCase interface {
	CreateSession(user *entity.User) string
	GetSession(uuid string) (*entity.User, error)
	DeleteSession(uuid string)
}
