package usecases

import (
	"belobetty-queue-manager/domain"
	"encoding/json"
)

type QueueProducerMock struct {
	queue []*domain.MessageDTO
}

func (q *QueueProducerMock) SendMessage(msg []byte) error {

	if q.queue == nil {
		q.queue = make([]*domain.MessageDTO, 0)
	}

	var entityDto domain.MessageDTO
	_ = json.Unmarshal(msg, &entityDto)

	q.queue = append(q.queue, &entityDto)
	return nil
}
