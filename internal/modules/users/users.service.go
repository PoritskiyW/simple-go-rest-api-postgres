package users

import (
	"errors"
	"strings"
)

func validateCreateUserRequest(req CreateUserRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validateUpdateUserRequest(req UpdateUserRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

func isValidEmail(email string) bool {
	// Simple email validation
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Service functions that add business logic
func GetAllUsersService() ([]User, error) {
	return GetAllUsers()
}

func GetUserByIDService(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return GetUserByID(id)
}

func CreateUserService(req CreateUserRequest) (*User, error) {
	if err := validateCreateUserRequest(req); err != nil {
		return nil, err
	}
	return CreateNewUser(req)
}

func UpdateUserService(id int, req UpdateUserRequest) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	if err := validateUpdateUserRequest(req); err != nil {
		return nil, err
	}
	return UpdateExistingUser(id, req)
}
