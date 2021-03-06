package entity

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID          ID
	Username    string
	Email       string `validate:"email"`
	FirstName   string
	LastName    string
	BirthDate   time.Time
	AboutMe     string
	PathToPhoto string
	Password    string `validate:"min=5"`
	IsPublic    bool
}

//NewUser ...
func NewUser(username, email, firstName, lastName, aboutMe, pathToPhoto, birthDate, password string) (*User, error) {
	user := &User{
		Username:    username,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		BirthDate:   StringToTime(birthDate),
		AboutMe:     aboutMe,
		PathToPhoto: pathToPhoto,
		Password:    password,
		IsPublic:    true,
	}
	user.EncryptPassword()
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

//Validate ...
func (u *User) Validate() error {
	var validate *validator.Validate = validator.New()
	return validate.Struct(u)
}

//EncryptPassword ...
func (u *User) EncryptPassword() {
	u.Password, _ = encrypt(u.Password)
}

//ComparePasswords ...
func (u *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

//TestUser ...
func TestUser(username string, password string) *User {
	return &User{
		Username:  username,
		Email:     username + "@gmail.com",
		FirstName: "First Name",
		LastName:  "Last Name",
		BirthDate: StringToTime("1994-09-18"),
		AboutMe:   "lorem ipsum vsyakaya hren liw by zapolnit eto pole",
		Password:  password,
		IsPublic:  true,
	}
}
