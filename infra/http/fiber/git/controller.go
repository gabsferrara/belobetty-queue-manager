package git

import (
	"belobetty-queue-manager/application/usecases"
	"belobetty-queue-manager/infra/queue"
)

type Controller struct {
	usecases.SenderUseCaseInterface
}

func NewGitController() (*Controller, error) {

	q, err := queue.NewRabbitMQ("git")
	if err != nil {
		return nil, err
	}

	return &Controller{usecases.NewSenderUseCase(q)}, nil
}
