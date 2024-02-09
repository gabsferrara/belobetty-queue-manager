package usecases

import (
	"belobetty-queue-manager/domain"
	"belobetty-queue-manager/infra/queue"
	"encoding/json"
)

type SenderUseCase struct {
	queue.Producer
}

type SenderUseCaseInterface interface {
	Exec(entity domain.Entity, action, functionality, user string) error
}

func NewSenderUseCase(sender queue.Producer) *SenderUseCase {
	return &SenderUseCase{sender}
}

func (s *SenderUseCase) Exec(entity domain.Entity, action, functionality, company string) error {
	message := &domain.MessageDTO{
		Company:       company,
		Action:        action,
		Functionality: functionality,
		Entity:        entity,
	}

	if action != "delete" {
		if err := entity.Validate(); err != nil {
			return err
		}
	}

	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := s.SendMessage(msg); err != nil {
		return err
	}

	return nil
}
