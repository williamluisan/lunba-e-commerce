package consumer

import (
	"context"
	"fmt"
	"time"

	repositoryMysqlImpl "lunba-e-commerce/internal/infrastructure/mysql/payment"
	conn "lunba-e-commerce/internal/transport/rabbitmq"
	service "lunba-e-commerce/internal/usecase/payment"
)

func PaymentConsume(conn *conn.RabbitMQConn) {
	paymentService := service.New(repositoryMysqlImpl.New())

	ctx := context.Background()
	
	/* dummy */
	msgs := consumeMock()
	for i := range msgs {
		orderId := int64(i)
		if _, err := paymentService.Process(ctx, orderId); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
	/* // */

	/* real rabbitmq */
	// ch := conn.AmqpChan

	// q, err := ch.QueueDeclare(
	// 	"payments", // name
	// 	false, 		// durable
	// 	false, 		// delete when unused
	// 	false, 		// exclusive
	// 	false, 		// no-wait
	// 	nil, 		//argument
	// )
	// conn.FailOnError(err, "Failed to declare a queue")

	// msgs, err := ch.Consume(
	// 	q.Name, // queue
	// 	"", 	// consumer
	// 	true, 	// auto-ack
	// 	false, 	// exclusive
	// 	false, 	// no-local
	// 	false, 	// no-wait
	// 	nil, 	// args
	// )

	// var forever chan struct{}

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)	
	// 	}
	// }()

	// <- forever
	/* // */
}

func consumeMock() <-chan int {
	ch := make(chan int)

	totalMsg := 50

	go func() {
		defer close(ch)
		for i := range totalMsg {
			ch <- i
		}
	}()

	return ch
}