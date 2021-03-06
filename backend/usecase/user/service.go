package user

import (
	"git.01.alem.school/Kusbek/social-network/backend/entity"
)

//Service ...
type Service struct {
	repo Repository
}

//NewService ...
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateUser ...
func (s *Service) CreateUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password string) (*entity.User, error) {
	u, err := entity.NewUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password)
	if err != nil {
		return nil, err
	}
	id, err := s.repo.Create(u)
	if err != nil {
		return nil, err
	}
	u.ID = id
	return u, nil
}

//FindUser ...
func (s *Service) FindUser(nickmail string) (*entity.User, error) {
	return s.repo.Find(nickmail)
}
