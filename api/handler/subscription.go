package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.01.alem.school/Kusbek/social-network/api/middleware"
	"git.01.alem.school/Kusbek/social-network/api/presenter"
	"git.01.alem.school/Kusbek/social-network/usecase/session"
	"git.01.alem.school/Kusbek/social-network/usecase/subscription"
)

// func getFollowers() {

// }

func follow(subscriptionService subscription.UseCase) http.HandlerFunc {
	var input struct {
		FollowingID int `json:"following_id"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		err = subscriptionService.Follow(r.Context().Value(middleware.UserID).(int), input.FollowingID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})
}

func unfollow(subscriptionService subscription.UseCase) http.HandlerFunc {
	var input struct {
		FollowingID int `json:"following_id"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		err = subscriptionService.Unfollow(r.Context().Value(middleware.UserID).(int), input.FollowingID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})
}

func isFollowing(subscriptionService subscription.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		followingID, err := strconv.Atoi(r.URL.Query().Get("following_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("following_id is a required parameter"))
			return
		}
		isFollowing, err := subscriptionService.IsFollowing(r.Context().Value(middleware.UserID).(int), followingID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"is_following": isFollowing,
		})
	})
}
func getFollowers(subscriptionService subscription.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		profileID, err := strconv.Atoi(r.URL.Query().Get("profile_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("profile_id is a required parameter"))
			return
		}
		followers, err := subscriptionService.GetFollowers(profileID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		followersJSON := make([]*presenter.User, 0, len(followers))
		for _, follower := range followers {
			followersJSON = append(followersJSON, &presenter.User{
				ID:          follower.ID,
				FirstName:   follower.FirstName,
				LastName:    follower.LastName,
				PathToPhoto: follower.PathToPhoto,
			})
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"followers_list": followersJSON,
		})
	})
}

func getFollowingUsers(subscriptionService subscription.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		profileID, err := strconv.Atoi(r.URL.Query().Get("profile_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("profile_id is a required parameter"))
			return
		}
		followingUsers, err := subscriptionService.GetFollowingUsers(profileID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		followingUserJSON := make([]*presenter.User, 0, len(followingUsers))
		for _, f := range followingUsers {
			followingUserJSON = append(followingUserJSON, &presenter.User{
				ID:          f.ID,
				FirstName:   f.FirstName,
				LastName:    f.LastName,
				PathToPhoto: f.PathToPhoto,
			})
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"following_list": followingUserJSON,
		})
	})
}

func MakeSubscriptionHandlers(r *http.ServeMux, sessionService session.UseCase, subscriptionService subscription.UseCase) {
	r.Handle("/api/follow", middleware.Auth(sessionService, follow(subscriptionService)))
	r.Handle("/api/unfollow", middleware.Auth(sessionService, unfollow(subscriptionService)))
	r.Handle("/api/isfollowing", middleware.Auth(sessionService, isFollowing(subscriptionService)))
	r.Handle("/api/followers", middleware.Auth(sessionService, getFollowers(subscriptionService)))
	r.Handle("/api/following", middleware.Auth(sessionService, getFollowingUsers(subscriptionService)))
}
