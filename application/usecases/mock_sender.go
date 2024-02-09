package usecases

import (
	"belobetty-queue-manager/domain"
	"encoding/json"
)

type QueueProducerMock struct {
	queue []*domain.MessageDTO
}

type EntityMock struct {
	Value string `json:"value"`
}

func (q *QueueProducerMock) SendMessage(msg []byte) error {
	if q.queue == nil {
		q.queue = make([]*domain.MessageDTO, 0)
	}

	var entity domain.MessageDTO
	_ = json.Unmarshal(msg, &entity)

	q.queue = append(q.queue, &entity)
	return nil
}

func (e EntityMock) Validate() error {
	return nil
}
