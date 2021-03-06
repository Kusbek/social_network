package subscription

import "git.01.alem.school/Kusbek/social-network/entity"

type Reader interface {
	IsFollowing(userID int, followingID int) (bool, error)
	GetFollowers(profileID int) ([]*entity.User, error)
	GetFollowingUsers(profileID int) ([]*entity.User, error)
	GetFollowRequests(profileID int) ([]*entity.User, error)
}

//Writer ...
type Writer interface {
	Follow(userID int, followingID int) error
	RequestFollow(userID int, followingID int) error
	Unfollow(userID int, followingID int) error
	AcceptFollowRequest(userID, followerID int) error
}

type Repository interface {
	Reader
	Writer
}

//UseCase ...
type UseCase interface {
	Follow(userID, followingID int) error
	Unfollow(userID, followingID int) error
	IsFollowing(userID int, followingID int) (bool, error)
	GetFollowers(profileID int) ([]*entity.User, error)
	GetFollowingUsers(profileID int) ([]*entity.User, error)
	GetFollowRequests(profileID int) ([]*entity.User, error)
	AcceptFollowRequest(userID, followerID int) error
}
