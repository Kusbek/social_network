package group

import "git.01.alem.school/Kusbek/social-network/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateGroup(ownerID entity.ID, title, description string) (*entity.Group, error) {
	g, err := entity.NewGroup(ownerID, title, description)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Create(g)
	if err != nil {
		return nil, err
	}

	g.ID = id
	return g, nil
}

func (s *Service) GetGroup(id entity.ID) (*entity.Group, error) {
	return s.repo.Get(id)
}

func (s *Service) GetGroups() ([]*entity.Group, error) {
	return s.repo.GetList()
}
