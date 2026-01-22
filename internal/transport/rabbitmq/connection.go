package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConn struct {
	AmqpChan	*amqp.Channel
}

func Conn() *RabbitMQConn {
	var rabbitMQConn RabbitMQConn

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
	rabbitMQConn.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	rabbitMQConn.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	return &RabbitMQConn{
		AmqpChan: ch,
	}
}

func (r *RabbitMQConn) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg ,err)
	}
}