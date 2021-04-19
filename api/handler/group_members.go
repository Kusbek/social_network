package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"git.01.alem.school/Kusbek/social-network/api/middleware"
	"git.01.alem.school/Kusbek/social-network/api/presenter"
	"git.01.alem.school/Kusbek/social-network/usecase/group"
	"git.01.alem.school/Kusbek/social-network/usecase/user"
)

func groupInviteHandlers(groupService group.UseCase, userService user.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			inviteUser(w, r, groupService, userService)
		case "PUT":
			acceptInvite(w, r, groupService)
		default:
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
		}
	})
}

func acceptInvite(w http.ResponseWriter, r *http.Request, groupService group.UseCase) {
	var input struct {
		GroupID int `json:"group_id"`
	}

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
}

func inviteUser(w http.ResponseWriter, r *http.Request, groupService group.UseCase, userService user.UseCase) {
	var input struct {
		GroupID  int    `json:"group_id"`
		Nickmail string `json:"nickmail"`
	}

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

	if g.OwnerID == u.ID {
		errorResponse(w, http.StatusBadRequest, fmt.Errorf("you can't invite yourself"))
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
