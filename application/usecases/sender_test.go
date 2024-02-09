package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSenderUseCase_Exec(t *testing.T) {
	mockSQueue := new(QueueProducerMock)
	s := NewSenderUseCase(mockSQueue)
	ent1 := EntityMock{Value: "Test1"}

	err := s.Exec(ent1, "acao1", "funcionalidade1", "empresa1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mockSQueue.queue))
	err = s.Exec(ent1, "acao2", "func2", "emp2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(mockSQueue.queue))

}
