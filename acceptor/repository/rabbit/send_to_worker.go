package rabbit

import (
	"github.com/streadway/amqp"
)

func SenderToWorker(ch *amqp.Channel, body []byte) error {

	// интересующая нас очередь
	queue, err := ch.QueueDeclare(
		"messages", // name название созданной очереди
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// отправим сообщение в очередь
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			Body: body,
		})
	if err != nil {
		return err
	}

	return nil
}
