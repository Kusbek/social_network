package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.01.alem.school/Kusbek/social-network/api/middleware"
	"git.01.alem.school/Kusbek/social-network/api/presenter"
	"git.01.alem.school/Kusbek/social-network/entity"
	"git.01.alem.school/Kusbek/social-network/usecase/session"
	"git.01.alem.school/Kusbek/social-network/usecase/subscription"
	"git.01.alem.school/Kusbek/social-network/usecase/user"
)

func signup(sessionService session.UseCase, userService user.UseCase) http.HandlerFunc {
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
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
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

func login(sessionService session.UseCase, userService user.UseCase) http.HandlerFunc {
	var input struct {
		Credentials string `json:"creds"`
		Password    string `json:"password"`
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

		u, err := userService.FindUser(input.Credentials)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, fmt.Errorf("user not found"))
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
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			errorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized, no cookie"))
			return
		}

		u, err := sessionService.GetSession(cookie.Value)
		if err != nil {
			errorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
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

func logout(sessionService session.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			successResponse(w, http.StatusOK, map[string]interface{}{
				"success": true,
			})
			return
		}
		deleteCookie(w, sessionService, cookie)
		successResponse(w, http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})
}

func getUser(userService user.UseCase, subService subscription.UseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}

		userID := r.Context().Value(middleware.UserID).(int)

		profileID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, fmt.Errorf("id is a required parameter"))
			return
		}

		u, err := userService.GetUser(profileID)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		if !u.IsPublic && userID != profileID {
			isFollowing, err := subService.IsFollowing(userID, profileID)
			if err != nil {
				errorResponse(w, http.StatusInternalServerError, err)
				return
			}

			if !isFollowing {
				userJSON := &presenter.User{
					IsPublic: u.IsPublic,
				}
				successResponse(w, http.StatusOK, userJSON)
				return
			}
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
			IsPublic:    u.IsPublic,
		}
		successResponse(w, http.StatusOK, userJSON)
	})
}

func deleteCookie(w http.ResponseWriter, sessionService session.UseCase, cookie *http.Cookie) {
	sessionService.DeleteSession(cookie.Value)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}

func setCookie(w http.ResponseWriter, service session.UseCase, user *entity.User) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   service.CreateSession(user),
		Expires: time.Now().Add(10 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
}

func setProfileVisibility(userService user.UseCase) http.HandlerFunc {
	var input struct {
		IsPublic bool `json:"is_public"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			errorResponse(w, http.StatusMethodNotAllowed, fmt.Errorf("wrong method"))
			return
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}

		affectedRows, err := userService.ChangeVisibility(r.Context().Value(middleware.UserID).(int), input.IsPublic)
		if err != nil {
			if err == sql.ErrNoRows {
				errorResponse(w, http.StatusNotFound, err)
				return
			}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusOK, map[string]interface{}{
			"affected_rows": affectedRows,
		})
	})
}

//MakeUserHandlers ...
func MakeUserHandlers(r *http.ServeMux, sessionService session.UseCase, userService user.UseCase, subService subscription.UseCase) {
	r.Handle("/api/signup", signup(sessionService, userService))
	r.Handle("/api/login", login(sessionService, userService))
	r.Handle("/api/auth", authenticate(sessionService))
	r.Handle("/api/logout", logout(sessionService))
	r.Handle("/api/user", middleware.Auth(sessionService, getUser(userService, subService)))
	r.Handle("/api/user/setprofilevisibility", middleware.Auth(sessionService, setProfileVisibility(userService)))
}
