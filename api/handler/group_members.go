package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.01.alem.school/Kusbek/social-network/api/middleware"
	"git.01.alem.school/Kusbek/social-network/api/presenter"
	"git.01.alem.school/Kusbek/social-network/usecase/group"
	"git.01.alem.school/Kusbek/social-network/usecase/user"
)

func inviteUser(groupService group.UseCase, userService user.UseCase) http.HandlerFunc {
	var input struct {
		GroupID  int    `json:"group_id"`
		Nickmail string `json:"nickmail"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		u, err := userService.FindUser(input.Nickmail)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		g, err := groupService.GetGroup(input.GroupID)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		fmt.Println(u, g)
		if g.OwnerID == u.ID {
			errorResponse(w, http.StatusTeapot, fmt.Errorf("you can't invite yourself"))
			return
		}
		err = groupService.CreateInvitedByGroupRequest(u.ID, g.ID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		successResponse(w, http.StatusCreated, map[string]interface{}{
			"success": true,
		})
	})
}

func acceptInvite(groupService group.UseCase) http.HandlerFunc {
	var input struct {
		GroupID int `json:"group_id"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}

		g, err := groupService.GetGroup(input.GroupID)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		err = groupService.AcceptInvite(r.Context().Value(middleware.UserID).(int), g.ID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusCreated, map[string]interface{}{
			"success": true,
		})
	})
}

func getGroupInvites(groupService group.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}

		groupInvites, err := groupService.GetInvites(r.Context().Value(middleware.UserID).(int))
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		groupInvitesJSON := make([]*presenter.Group, 0, len(groupInvites))
		for _, g := range groupInvites {
			groupInvitesJSON = append(groupInvitesJSON, &presenter.Group{
				ID:          g.ID,
				Title:       g.Title,
				Description: g.Description,
			})
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"group_invites": groupInvitesJSON,
		})
	})
}

func getGroupMembers(groupService group.UseCase, userService user.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("group_id is a required parameter"))
			return
		}

		members, err := groupService.GetGroupMembers(groupID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		membersJSON := make([]*presenter.User, 0, len(members))
		for _, m := range members {
			membersJSON = append(membersJSON, &presenter.User{
				ID:          m.ID,
				FirstName:   m.FirstName,
				LastName:    m.LastName,
				PathToPhoto: m.PathToPhoto,
			})
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"group_members": membersJSON,
		})
	})
}

func isGroupMember(groupService group.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("group_id is a required parameter"))
			return
		}
		userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("user_id is a required parameter"))
			return
		}

		isMember, err := groupService.IsGroupMember(userID, groupID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		if isMember {
			successResponse(w, http.StatusOK, map[string]interface{}{
				"is_group_member":    true,
				"request_is_pending": false,
			})
			return
		}

		isPending, err := groupService.RequestIsPending(userID, groupID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"is_group_member":    false,
			"request_is_pending": isPending,
		})
	})
}

func requestToJoin(groupService group.UseCase) http.HandlerFunc {
	var input struct {
		GroupID int `json:"group_id"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		g, err := groupService.GetGroup(input.GroupID)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		err = groupService.CreateJoinGroupRequest(r.Context().Value(middleware.UserID).(int), g.ID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		successResponse(w, http.StatusCreated, map[string]interface{}{
			"success": true,
		})
	})
}
