package git

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	Owner        = "OWNER"
	Collaborator = "COLLABORATOR"
	Maintainer   = "MAINTAINER"
	Reader       = "READER"
)

type RepositoryGit struct {
	ID          int64             `json:"ID"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Private     bool              `json:"private,omitempty"`
	Users       UserRepositoryGit `json:"users,omitempty"`
}

type UserRepositoryGit map[string]string

func NewRepository(name, description string, private bool, users UserRepositoryGit) (*RepositoryGit, error) {
	repository := &RepositoryGit{
		Name:        name,
		Description: description,
		Private:     private,
		Users:       users,
	}

	err := repository.Validate()
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (u *RepositoryGit) Validate() error {
	err := validRepositoryName(u.Name)
	if err != nil {
		return err
	}
	if len(u.Description) > 350 {
		return errors.New("description cannot be more than 350 characters")
	}
	err = validPermissionCategory(u.Users)
	if err != nil {
		return err
	}
	return nil
}

func validPermissionCategory(users UserRepositoryGit) error {
	if users == nil {
		return nil
	}
	result := "invalid user category permission:"
	for user, permission := range users {
		switch strings.ToUpper(permission) {
		case Owner, Collaborator, Maintainer, Reader:
			continue
		default:
			result += fmt.Sprintf("\nUser: %s with permission: %s", user, permission)
		}
	}
	if result != "invalid user category permission:" {
		return errors.New(result)
	}
	return nil
}

func validRepositoryName(n string) error {
	if n == "" {
		return errors.New("n repository is required")
	}
	if len(n) > 100 {
		return errors.New("name repository is too long (maximum is 100 characters)")
	}
	regexPattern := "^[a-zA-Z0-9._-]+$"
	ok, err := regexp.MatchString(regexPattern, n)
	if err != nil || !ok {
		return errors.New("the repository name can only contain ASCII letters, digits, and the characters ., -, and _")
	}
	return nil
}
