package handler

import (
	"net/http"

	"git.01.alem.school/Kusbek/social-network/backend/api/presenter"
	"git.01.alem.school/Kusbek/social-network/backend/usecase/user"
	"github.com/gin-gonic/gin"
)

func createUser(service user.UseCase) gin.HandlerFunc {
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
	return gin.HandlerFunc(func(c *gin.Context) {
		err := c.ShouldBindJSON(&input)
		if err != nil {
			errorResponse(c, http.StatusBadRequest, err)
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
			input.Password)

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

		successResponse(c, http.StatusOK, userJSON)
	})
}

func MakeUserHandlers(r *gin.Engine, service user.UseCase) {
	r.Handle(http.MethodPost, "/signup", createUser(service))
}
