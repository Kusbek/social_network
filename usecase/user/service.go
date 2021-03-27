package user

import (
	"git.01.alem.school/Kusbek/social-network/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

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

func (s *Service) FindUser(nickmail string) (*entity.User, error) {
	return s.repo.Find(nickmail)
}

func (s *Service) GetUser(id int) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s *Service) ChangeVisibility(id int, isPublic bool) (int, error) {
	if isPublic {
		return s.repo.ChangeVisibility(id, 1)
	}
	return s.repo.ChangeVisibility(id, 0)
}
