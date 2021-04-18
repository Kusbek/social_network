package presenter

import (
	"git.01.alem.school/Kusbek/social-network/entity"
)

type Group struct {
	ID          entity.ID `json:"id"`
	Owner       *User     `json:"owner,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
