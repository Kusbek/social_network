package repository

import (
	"database/sql"

	"git.01.alem.school/Kusbek/social-network/entity"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) Follow(userID, followingID int) error {
	_, err := r.db.Exec(`INSERT INTO followers (user_id, following_id) VALUES($1,$2)`, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) Unfollow(userID, followingID int) error {
	_, err := r.db.Exec(`DELETE FROM followers WHERE user_id=$1 AND following_id=$2`, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) IsFollowing(userID, followingID int) (bool, error) {
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM followers WHERE user_id=$1 AND following_id=$2)`, userID, followingID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SubscriptionRepository) GetFollowers(profileID int) ([]*entity.User, error) {
	rows, err := r.db.Query(`SELECT id, first_name, last_name from users WHERE id IN (SELECT following_id from followers WHERE user_id=$1)`, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	followers := make([]*entity.User, 0)
	for rows.Next() {
		follower := &entity.User{}
		err = rows.Scan(
			&follower.ID,
			&follower.FirstName,
			&follower.LastName,
		)
		if err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}
