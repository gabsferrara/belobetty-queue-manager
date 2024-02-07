package git_test

import (
	"belobetty-queue-manager/domain/git"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const initialFilePathUser = "../../docs/git/users/"

func TestCreateUser(t *testing.T) {
	user, err := git.NewUser("gabGit", "gabriel@belobetty.com", "Gabriel Ferrara", nil)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, "gabgit", user.UserName)
	assert.Equal(t, user.Email, "gabriel@belobetty.com")

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "ok.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)

	err = userJson.Validate()
	assert.Nil(t, err)
	assert.Equal(t, "gabGit", userJson.UserName)
	assert.Equal(t, "gabriel@belobetty.com", userJson.Email)
	assert.Equal(t, 2, len(userJson.Permissions))
	assert.Equal(t, "Owner", userJson.Permissions[0].Permission)
}

func TestCreateUserWithOutUserName(t *testing.T) {
	user, err := git.NewUser("", "gabriel@belobetty.com", "Gabriel Ferrara", nil)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withOutUserName.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "username is required", err.Error())
}

func TestCreateUserInvalidUserName(t *testing.T) {
	user, err := git.NewUser("StringToNameWithMoreThan34Caracteres12345", "gabriel@belobetty.com", "Gabriel Ferrara", nil)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withInvalidUserNameBiggestThan39.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "username is too long (maximum is 39 characters)", err.Error())

	jsonData, _ = os.ReadFile(initialFilePathUser + "withInvalidUserName.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid username, may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen", err.Error())
}

func TestCreateUserWithOutOrInvalidEmail(t *testing.T) {
	user, err := git.NewUser("gabGit", "gabriel@.com", "Gabriel Ferrara", nil)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withInvalidEmail.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid e-mail", err.Error())
}

func TestCreateUserWithOutEmail(t *testing.T) {
	user, err := git.NewUser("gabGit", "", "Gabriel Ferrara", nil)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withOutEmail.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestCreateUserWithOutName(t *testing.T) {
	user, err := git.NewUser("gabGit", "gabriel@belobetty.com", "", nil)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withOutName.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestCreateUserWithInvalidPermission(t *testing.T) {

	access, err := git.NewAccessRepository("RepoTest", "Error")
	assert.Nil(t, access)
	assert.NotNil(t, err)

	var userJson = new(git.UserGit)

	jsonData, _ := os.ReadFile(initialFilePathUser + "withInvalidPermissionInRepository.json")
	err = json.Unmarshal(jsonData, userJson)
	assert.Nil(t, err)
	assert.NotNil(t, userJson)
	err = userJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid user category permission: PermissionInvalid", err.Error())
}
