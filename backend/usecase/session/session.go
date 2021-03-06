package session

import "git.01.alem.school/Kusbek/social-network/backend/entity"

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

//CreateSession ...
func (s *Service) CreateSession(user *entity.User) string {
	return s.repo.Create(user)
}

//GetSession ...
func (s *Service) GetSession(uuid string) (*entity.User, error) {
	return s.repo.Get(uuid)
}

//DeleteSession ...
func (s *Service) DeleteSession(uuid string) {
	s.repo.Delete(uuid)
}
