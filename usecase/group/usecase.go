package group

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

func (s *Service) CreateGroup(ownerID entity.ID, title, description string) (*entity.Group, error) {
	g, err := entity.NewGroup(ownerID, title, description)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Create(g)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateInvitedByGroupRequest(g.OwnerID, id)
	if err != nil {
		return nil, err
	}

	err = s.repo.AcceptInvite(g.OwnerID, id)
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

func (s *Service) GetInvites(userID entity.ID) ([]*entity.Group, error) {
	return s.repo.GetInvites(userID)
}

func (s *Service) AcceptInvite(userID, groupID entity.ID) error {
	return s.repo.AcceptInvite(userID, groupID)
}

func (s *Service) CreateInvitedByGroupRequest(userID, groupID entity.ID) error {
	return s.repo.CreateInvitedByGroupRequest(userID, groupID)
}

func (s *Service) CreateJoinGroupRequest(userID, groupID entity.ID) error {
	return s.repo.CreateJoinGroupRequest(userID, groupID)
}

func (s *Service) GetGroupMembers(groupID entity.ID) ([]*entity.User, error) {
	return s.repo.GetGroupMembers(groupID)
}

func (s *Service) IsGroupMember(userID, groupID entity.ID) (bool, error) {
	return s.repo.IsGroupMember(userID, groupID)
}

func (s *Service) RequestIsPending(userID, groupID entity.ID) (bool, error) {
	return s.repo.RequestIsPending(userID, groupID)
}
