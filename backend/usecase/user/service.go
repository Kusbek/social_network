package user

import (
	"git.01.alem.school/Kusbek/social-network/backend/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password string) (entity.ID, error) {
	u, err := entity.NewUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birhDate, password)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(u)
}

func (s *Service) FindUser(nickmail string, password string) (*entity.User, error) {
	return s.repo.Find(nickmail)
}
