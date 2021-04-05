package entity

import "fmt"

type Group struct {
	ID          ID
	OwnerID     ID
	Title       string
	Description string
}

func NewGroup(ownerID ID, title, description string) (*Group, error) {
	group := &Group{
		OwnerID:     ownerID,
		Title:       title,
		Description: description,
	}
	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *Group) Validate() error {
	errMsg := "%v should not be empty"
	if g.Title == "" {
		return fmt.Errorf(errMsg, "title")
	}

	if g.Description == "" {
		return fmt.Errorf(errMsg, "description")
	}

	return nil
}

func TestGroup(ownerID ID) *Group {
	return &Group{
		OwnerID:     ownerID,
		Title:       "Test group title",
		Description: "Test group description",
	}
}
