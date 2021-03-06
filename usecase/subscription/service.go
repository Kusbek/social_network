package subscription

import (
	"git.01.alem.school/Kusbek/social-network/entity"
	"git.01.alem.school/Kusbek/social-network/usecase/user"
)

type Service struct {
	userService user.UseCase
	repo        Repository
}

func NewService(r Repository, userService user.UseCase) *Service {
	return &Service{
		userService: userService,
		repo:        r,
	}
}

func (s *Service) Follow(userID, followingID int) error {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return err
	}
	if !followingUser.IsPublic {
		return s.repo.RequestFollow(userID, followingUser.ID)
	}
	return s.repo.Follow(userID, followingUser.ID)
}

func (s *Service) AcceptFollowRequest(userID, followerID int) error {
	followerUser, err := s.userService.GetUser(followerID)
	if err != nil {
		return err
	}

	return s.repo.AcceptFollowRequest(userID, followerUser.ID)
}

func (s *Service) IsFollowing(userID, followingID int) (bool, error) {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return false, err
	}
	return s.repo.IsFollowing(userID, followingUser.ID)
}

func (s *Service) Unfollow(userID, followingID int) error {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return err
	}
	return s.repo.Unfollow(userID, followingUser.ID)
}

func (s *Service) GetFollowers(profileID int) ([]*entity.User, error) {
	return s.repo.GetFollowers(profileID)
}

func (s *Service) GetFollowRequests(profileID int) ([]*entity.User, error) {
	return s.repo.GetFollowRequests(profileID)
}

func (s *Service) GetFollowingUsers(profileID int) ([]*entity.User, error) {
	return s.repo.GetFollowingUsers(profileID)
}
