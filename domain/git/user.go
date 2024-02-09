package git

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

type UserGit struct {
	UserName    string                    `json:"username"`
	Email       string                    `json:"email,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Permissions []AccessRepositoryUserGit `json:"permissions,omitempty"`
}

type AccessRepositoryUserGit struct {
	Repository string `json:"repository,omitempty"`
	Permission string `json:"permission,omitempty"`
}

func NewUser(userName, email, name string, permission []AccessRepositoryUserGit) (*UserGit, error) {
	user := &UserGit{
		UserName:    strings.ToLower(userName),
		Email:       email,
		Name:        name,
		Permissions: permission,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAccessRepository(repositoryName, permission string) (*AccessRepositoryUserGit, error) {
	access := &AccessRepositoryUserGit{
		Repository: strings.ToLower(repositoryName),
		Permission: strings.ToLower(permission),
	}
	err := access.validate()
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (r *AccessRepositoryUserGit) validate() error {
	switch strings.ToUpper(r.Permission) {
	case Owner, Collaborator, Maintainer, Reader:
		return nil
	default:
		return fmt.Errorf("invalid user category permission: %s", r.Permission)
	}

}

func (u *UserGit) Validate() error {
	err := validUserName(u.UserName)
	if err != nil {
		return err
	}
	err = validEmail(u.Email)
	if err != nil {
		return err
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	for _, access := range u.Permissions {
		err = access.validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func validUserName(userName string) error {
	if userName == "" {
		return errors.New("username is required")
	}
	if len(userName) > 39 {
		return errors.New("username is too long (maximum is 39 characters)")
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)
	if !regex.MatchString(userName) {
		return errors.New("invalid username, may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen")
	}
	return nil
}

func validEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid e-mail")
	}

	return nil
}
