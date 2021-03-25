package repository

import "database/sql"

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) Follow(userID, followingID int) error {
	_, err := r.db.Exec(
		`
		INSERT INTO followers (user_id, following_id)
		VALUES($1,$2)
		`, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) Unfollow(userID, followingID int) error {
	_, err := r.db.Exec(
		`
		DELETE FROM followers WHERE user_id=$1 AND following_id=$2
		`, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) IsFollowing(userID, followingID int) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`
		SELECT EXISTS(SELECT 1 FROM followers WHERE user_id=$1 AND following_id=$2)
		`, userID, followingID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
