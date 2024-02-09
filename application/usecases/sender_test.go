package usecases

import (
	"belobetty-queue-manager/domain/git"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSenderUseCase_Exec(t *testing.T) {
	mockSQueue := new(QueueProducerMock)
	s := NewSenderUseCase(mockSQueue)

	var ent1 = new(git.UserGit)

	jsonData, _ := os.ReadFile("../../docs/git/users/ok.json")
	err := json.Unmarshal(jsonData, ent1)
	assert.Nil(t, err)
	assert.NotNil(t, ent1)

	err = s.Exec(ent1, "anyAction", "funcionalidade1", "empresa1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mockSQueue.queue))

	ent2 := new(git.RepositoryGit)
	jsonData, _ = os.ReadFile("../../docs/git/repositories/ok.json")
	err = json.Unmarshal(jsonData, ent2)
	err = s.Exec(ent1, "anyAction", "func2", "emp2")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(mockSQueue.queue))

}

func TestSenderUseCase_ExecWhenErrorAndActionNotDelete(t *testing.T) {
	mockSQueue := new(QueueProducerMock)
	s := NewSenderUseCase(mockSQueue)

	ent := new(git.RepositoryGit)
	jsonData, _ := os.ReadFile("../../docs/git/repositories/invalidUserPermission.json")
	err := json.Unmarshal(jsonData, ent)

	err = s.Exec(ent, "anyAction", "funcionalidade1", "empresa1")
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(mockSQueue.queue))

}
