package presenter

import "git.01.alem.school/Kusbek/social-network/entity"

type Group struct {
	ID          entity.ID `json:"id"`
	OwnerID     entity.ID `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
