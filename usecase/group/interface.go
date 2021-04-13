package group

import "git.01.alem.school/Kusbek/social-network/entity"

type Reader interface {
	Get(id entity.ID) (*entity.Group, error)
	GetList() ([]*entity.Group, error)
}

//Writer ...
type Writer interface {
	Create(group *entity.Group) (entity.ID, error)
	CreateInvitedByGroupRequest(userID, groupID entity.ID) error
}

//Repository ...
type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateGroup(ownerID entity.ID, title, description string) (*entity.Group, error)
	GetGroup(id entity.ID) (*entity.Group, error)
	GetGroups() ([]*entity.Group, error)
	CreateInvitedByGroupRequest(userID, groupID entity.ID) error
}
