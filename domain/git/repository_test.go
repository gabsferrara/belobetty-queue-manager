package git_test

import (
	"belobetty-queue-manager/domain/git"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const initialFilePathRepo = "../../docs/git/repositories/"

func TestCreateRepository(t *testing.T) {
	users := make(git.UserRepositoryGit)
	users["Gabriel"] = "Owner"
	users["Dani"] = "collaborator"
	repo, err := git.NewRepository("Repo_name", "Description Test", true, users)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, repo.Name, "Repo_name")

	repoJson := new(git.RepositoryGit)
	jsonData, _ := os.ReadFile(initialFilePathRepo + "ok.json")
	err = json.Unmarshal(jsonData, repoJson)
	assert.Nil(t, err)
	assert.NotNil(t, repoJson)

	err = repoJson.Validate()
	assert.Nil(t, err)

	assert.Equal(t, "my_repo", repoJson.Name)
	assert.Equal(t, "OWNER", repoJson.Users["gabs_git"])
	assert.Equal(t, 4, len(repoJson.Users))

}

func TestCreateRepositoryWithInvalidName(t *testing.T) {
	repo, err := git.NewRepository("test@", "Description test", true, nil)
	assert.Nil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, "the repository name can only contain ASCII letters, digits, and the characters ., -, and _", err.Error())

	repoJson := new(git.RepositoryGit)
	jsonData, _ := os.ReadFile(initialFilePathRepo + "invalidName.json")
	err = json.Unmarshal(jsonData, repoJson)
	assert.Nil(t, err)
	assert.NotNil(t, repoJson)

	err = repoJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "the repository name can only contain ASCII letters, digits, and the characters ., -, and _", err.Error())

	jsonData, _ = os.ReadFile(initialFilePathRepo + "invalidNameMostThan100.json")
	err = json.Unmarshal(jsonData, repoJson)
	assert.Nil(t, err)
	assert.NotNil(t, repoJson)

	err = repoJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "name repository is too long (maximum is 100 characters)", err.Error())

}

func TestCreateRepositoryWithInvalidDescription(t *testing.T) {
	bigString := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" +
		"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"

	repo, err := git.NewRepository("Test", bigString, true, nil)
	assert.True(t, len(bigString) > 350)
	assert.Nil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "description cannot be more than 350 characters")

	repoJson := new(git.RepositoryGit)
	jsonData, _ := os.ReadFile(initialFilePathRepo + "invalidDescription.json")
	err = json.Unmarshal(jsonData, repoJson)
	assert.Nil(t, err)
	assert.NotNil(t, repoJson)

	err = repoJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "description cannot be more than 350 characters", err.Error())
}

func TestCreateRepositoryWithInvalidUserPermission(t *testing.T) {
	users := map[string]string{
		"UserOk":    "Maintainer",
		"UserFail1": "Master",
		"UserOk2":   "Collaborator",
	}

	repo, err := git.NewRepository("Repo_name", "Description Test", true, users)
	assert.Nil(t, repo)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid user category permission:\nUser: UserFail1 with permission: Master")

	repoJson := new(git.RepositoryGit)
	jsonData, _ := os.ReadFile(initialFilePathRepo + "invalidUserPermission.json")
	err = json.Unmarshal(jsonData, repoJson)
	assert.Nil(t, err)
	assert.NotNil(t, repoJson)

	err = repoJson.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "invalid user category permission:\nUser: bob_smith with permission: INVALID", err.Error())
}
