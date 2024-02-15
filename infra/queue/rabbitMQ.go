package queue

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

type RabbitMQProducer struct {
	con     *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func init() {
	_ = godotenv.Load()
	if envURL, ok := os.LookupEnv("RABBITMQ_URL"); ok {
		url = envURL
	}
}

var url = "guest@localhost:5672"

func NewRabbitMQ(routingKey string) (*RabbitMQProducer, error) {
	conn, err := amqp.Dial("amqp://guest:" + url)
	if err != nil {
		return nil, fmt.Errorf("failed to connecto to RabbitMQ: %s", err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %s", err.Error())
	}

	q, err := ch.QueueDeclare(
		routingKey, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %s", err.Error())
	}
	return &RabbitMQProducer{
		con:     conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (r *RabbitMQProducer) SendMessage(msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer r.con.Close()
	defer r.channel.Close()
	defer cancel()

	err := r.channel.PublishWithContext(ctx, "", r.queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})
	if err != nil {
		return err
	}

	return nil
}
