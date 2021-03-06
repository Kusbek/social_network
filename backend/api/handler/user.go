package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"git.01.alem.school/Kusbek/social-network/backend/api/middleware"
	"git.01.alem.school/Kusbek/social-network/backend/api/presenter"
	"git.01.alem.school/Kusbek/social-network/backend/entity"
	"git.01.alem.school/Kusbek/social-network/backend/usecase/session"
	"git.01.alem.school/Kusbek/social-network/backend/usecase/user"
)

func createUser(sessionService session.UseCase, userService user.UseCase) http.HandlerFunc {
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
		u, err := userService.CreateUser(
			input.Username,
			input.Email,
			input.FirstName,
			input.LastName,
			input.AboutMe,
			input.PathToPhoto,
			input.BirthDate,
			input.Password,
		)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		// userJSON := &presenter.User{
		// 	ID:          u.ID,
		// 	Username:    input.Username,
		// 	Email:       input.Email,
		// 	FirstName:   input.FirstName,
		// 	LastName:    input.LastName,
		// 	AboutMe:     input.AboutMe,
		// 	PathToPhoto: input.PathToPhoto,
		// 	BirthDate:   input.BirthDate,
		// }

		setCookie(w, sessionService, u)
		successResponse(w, http.StatusOK, map[string]bool{"success": true})
	})
}

func authorizeUser(sessionService session.UseCase, userService user.UseCase) http.HandlerFunc {
	var input struct {
		Credentials string `json:"creds"`
		Password    string `json:"password"`
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

		u, err := userService.FindUser(input.Credentials)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, fmt.Errorf("User not found"))
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		if err = u.ComparePasswords(input.Password); err != nil {
			errorResponse(w, http.StatusForbidden, err)
			return
		}

		// userJSON := &presenter.User{
		// 	ID:          u.ID,
		// 	Username:    u.Username,
		// 	Email:       u.Email,
		// 	FirstName:   u.FirstName,
		// 	LastName:    u.LastName,
		// 	AboutMe:     u.AboutMe,
		// 	PathToPhoto: u.PathToPhoto,
		// 	BirthDate:   entity.TimeToString(u.BirthDate),
		// }
		setCookie(w, sessionService, u)
		successResponse(w, http.StatusOK, map[string]bool{"success": true})
	})
}

func authenticate(sessionService session.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("Wrong Method"))
			return
		}
		cookie, err := r.Cookie("session_id")
		fmt.Println("Cookies: ", cookie.Value)
		if err == http.ErrNoCookie {
			errorResponse(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized, No Cookie"))
			return
		}

		u, err := sessionService.GetSession(cookie.Value)
		if err != nil {
			errorResponse(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
			return
		}

		userJSON := &presenter.User{
			ID:          u.ID,
			Username:    u.Username,
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			AboutMe:     u.AboutMe,
			PathToPhoto: u.PathToPhoto,
			BirthDate:   entity.TimeToString(u.BirthDate),
		}
		successResponse(w, http.StatusOK, userJSON)
	})
}

func setCookie(w http.ResponseWriter, service session.UseCase, user *entity.User) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   service.CreateSession(user),
		Expires: time.Now().Add(10 * time.Hour),
		Path:    "/",
	}
	fmt.Println(cookie.Value, user)
	http.SetCookie(w, &cookie)
}

//MakeUserHandlers ...
func MakeUserHandlers(r *http.ServeMux, sessionService session.UseCase, userService user.UseCase) {
	r.Handle("/signup", middleware.Cors(createUser(sessionService, userService)))
	r.Handle("/login", middleware.Cors(authorizeUser(sessionService, userService)))
	r.Handle("/auth", middleware.Cors(authenticate(sessionService)))
}
