package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"git.01.alem.school/Kusbek/social-network/backend/api/presenter"
	"git.01.alem.school/Kusbek/social-network/backend/entity"
	"git.01.alem.school/Kusbek/social-network/backend/usecase/user"
)

func createUser(service user.UseCase) http.Handler {
	var input struct {
		Username    string `json:"username,omitempty"`
		Email       string `json:"email,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		BirthDate   string `json:"birth_date,omitempty"`
		AboutMe     string `json:"about_me,omitempty"`
		PathToPhoto string `json:"path_to_photo,omitempty"`
		Password    string `json:"password,omitempty"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("Wrong Method"))
			return
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		id, err := service.CreateUser(
			input.Username,
			input.Email,
			input.FirstName,
			input.LastName,
			input.AboutMe,
			input.PathToPhoto,
			input.BirthDate,
			input.Password,
		)

		userJSON := &presenter.User{
			ID:          id,
			Username:    input.Username,
			Email:       input.Email,
			FirstName:   input.FirstName,
			LastName:    input.LastName,
			AboutMe:     input.AboutMe,
			PathToPhoto: input.PathToPhoto,
			BirthDate:   input.BirthDate,
		}

		successResponse(w, http.StatusOK, userJSON)
	})
}

func authorizeUser(service user.UseCase) http.HandlerFunc {
	var input struct {
		Credentials string `json:"creds"`
		Password    string `json:"password"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}

		user, err := service.FindUser(input.Credentials)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, fmt.Errorf("User not found"))
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		if err = user.ComparePasswords(input.Password); err != nil {
			errorResponse(w, http.StatusForbidden, err)
			return
		}

		userJSON := &presenter.User{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			AboutMe:     user.AboutMe,
			PathToPhoto: user.PathToPhoto,
			BirthDate:   entity.TimeToString(user.BirthDate),
		}

		successResponse(w, http.StatusOK, userJSON)
	})
}

//MakeUserHandlers ...
func MakeUserHandlers(r *http.ServeMux, service user.UseCase) {
	r.Handle("/signup", createUser(service))
	r.Handle("/login", authorizeUser(service))

}
