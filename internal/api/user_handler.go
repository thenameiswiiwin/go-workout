package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/thenameiswiiwin/go-workout/internal/store"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

func (h *UserHandler) validateRegisterRequest(req *registerUserRequest) error {
	if req.Username == "" {
		return errors.New("username is required")
	}

	if len(req.Username) > 50 {
		return errors.New("username must be less than 50 characters")
	}

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,50}$`)
	if !usernameRegex.MatchString(req.Username) {
		return errors.New("username can only contain letters, numbers, underscores, and hyphens, and must be between 3 and 50 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if err := validatePassword(req.Password); err != nil {
		return err
	}

	if len(req.Bio) > 160 {
		return errors.New("bio must be less than 160 characters")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasLower := regexp.MustCompile(`[a-z]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasDigit := regexp.MustCompile(`\d`)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`)

	switch {
	case !hasLower.MatchString(password):
		return errors.New("password must contain at least one lowercase letter")
	case !hasUpper.MatchString(password):
		return errors.New("password must contain at least one uppercase letter")
	case !hasDigit.MatchString(password):
		return errors.New("password must contain at least one digit")
	case !hasSpecial.MatchString(password):
		return errors.New("password must contain at least one special character (@$!%*?&)")
	}

	return nil
}
