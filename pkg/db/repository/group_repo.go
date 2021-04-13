package repository

import (
	"database/sql"

	"git.01.alem.school/Kusbek/social-network/entity"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) Create(group *entity.Group) (entity.ID, error) {
	res, err := r.db.Exec(`INSERT INTO groups (owner_id, title, description) VALUES($1,$2,$3)`, group.OwnerID, group.Title, group.Description)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *GroupRepository) CreateInvitedByGroupRequest(userID, groupID entity.ID) error {
	_, err := r.db.Exec(`INSERT INTO group_list (user_id, group_id, group_requested) VALUES($1,$2,1)`, userID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupRepository) Get(id entity.ID) (*entity.Group, error) {
	group := new(entity.Group)
	err := r.db.QueryRow(`SELECT id, owner_id, title, description FROM groups WHERE id=$1`, id).Scan(
		&group.ID,
		&group.OwnerID,
		&group.Title,
		&group.Description,
	)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepository) GetList() ([]*entity.Group, error) {

	rows, err := r.db.Query(`SELECT id, owner_id, title, description FROM groups`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	groups := make([]*entity.Group, 0)
	for rows.Next() {
		group := new(entity.Group)
		err = rows.Scan(
			&group.ID,
			&group.OwnerID,
			&group.Title,
			&group.Description,
		)

		if err != nil {
			return nil, err
		}

		groups = append(groups, group)

	}

	if err != nil {
		return nil, err
	}
	return groups, nil
}
