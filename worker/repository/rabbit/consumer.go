package rabbit

import (
	"github.com/streadway/amqp"
	"worker/repository/postgres"
)

func ConsumerMsg(ch *amqp.Channel, pgxClient postgres.PgxClient) (<-chan amqp.Delivery, error) {

	// интересующая нас очередь
	queue, err := ch.QueueDeclare(
		"messages", // name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// получим сообщения из очереди
	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
