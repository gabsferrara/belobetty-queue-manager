package fiber

import (
	"belobetty-queue-manager/application/usecases"
	"belobetty-queue-manager/infra/queue"
	"github.com/joho/godotenv"
	"os"
)

type Controller struct {
	usecases.SenderUseCaseInterface
}

func init() {
	_ = godotenv.Load()
	if routingKeyEnv, ok := os.LookupEnv("RABBITMQ_ROUTE_KEY_GIT"); ok {
		routingKey = routingKeyEnv
	}
}

var routingKey = "git"

func NewGitController() (*Controller, error) {

	q, err := queue.NewRabbitMQ(routingKey)
	if err != nil {
		return nil, err
	}

	return &Controller{usecases.NewSenderUseCase(q)}, nil
}
