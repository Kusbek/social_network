package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.01.alem.school/Kusbek/social-network/api/middleware"
	"git.01.alem.school/Kusbek/social-network/api/presenter"
	"git.01.alem.school/Kusbek/social-network/usecase/group"
	"git.01.alem.school/Kusbek/social-network/usecase/session"
)

func createGroup(w http.ResponseWriter, r *http.Request, groupService group.UseCase) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if r.Method != "POST" {
		errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	group, err := groupService.CreateGroup(r.Context().Value(middleware.UserID).(int), input.Title, input.Description)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
	}

	successResponse(w, http.StatusOK, group)

}

func getGroups(groupService group.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}

		groups, err := groupService.GetGroups()
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		groupsJSON := make([]*presenter.Group, 0, len(groups))
		for _, group := range groups {
			groupsJSON = append(groupsJSON, &presenter.Group{
				ID:          group.ID,
				OwnerID:     group.OwnerID,
				Title:       group.Title,
				Description: group.Description,
			})
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"group_list": groupsJSON,
		})
	})
}

func getGroup(w http.ResponseWriter, r *http.Request, groupService group.UseCase) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, fmt.Errorf("group_id is a required parameter"))
		return
	}

	group, err := groupService.GetGroup(groupID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(group)
	successResponse(w, http.StatusOK, &presenter.Group{
		ID:          group.ID,
		OwnerID:     group.OwnerID,
		Title:       group.Title,
		Description: group.Description,
	})
}

func groupHandlers(groupService group.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getGroup(w, r, groupService)
		case "POST":
			createGroup(w, r, groupService)
		default:
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
		}
	})
}

func MakeGroupHandlers(r *http.ServeMux, sessionService session.UseCase, groupService group.UseCase) {
	r.Handle("/api/group", middleware.Auth(sessionService, groupHandlers(groupService)))
	r.Handle("/api/groups", getGroups(groupService))
}
