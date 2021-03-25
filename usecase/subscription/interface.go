package subscription

type Reader interface {
	IsFollowing(userID int, followingID int) (bool, error)
}

//Writer ...
type Writer interface {
	Follow(userID int, followingID int) error
	Unfollow(userID int, followingID int) error
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
}
