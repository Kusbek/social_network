package presenter

import "git.01.alem.school/Kusbek/social-network/entity"

//User ...
type User struct {
	ID          entity.ID `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	BirthDate   string    `json:"birth_date,omitempty"`
	AboutMe     string    `json:"about_me,omitempty"`
	PathToPhoto string    `json:"path_to_photo,omitempty"`
}
