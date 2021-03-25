package subscription

import "git.01.alem.school/Kusbek/social-network/usecase/user"

//Service ...
type Service struct {
	userService user.UseCase
	repo        Repository
}

//NewService ...
func NewService(r Repository, userService user.UseCase) *Service {
	return &Service{
		userService: userService,
		repo:        r,
	}
}

//Follow ...
func (s *Service) Follow(userID, followingID int) error {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return err
	}
	return s.repo.Follow(userID, followingUser.ID)
}

//Follow ...
func (s *Service) IsFollowing(userID, followingID int) (bool, error) {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return false, err
	}
	return s.repo.IsFollowing(userID, followingUser.ID)
}

//Follow ...
func (s *Service) Unfollow(userID, followingID int) error {
	followingUser, err := s.userService.GetUser(followingID)
	if err != nil {
		return err
	}
	return s.repo.Unfollow(userID, followingUser.ID)
}
