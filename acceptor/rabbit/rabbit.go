package rabbit

import (
	"acceptor/config"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type RbtClient struct {
	ConnRbtMsg *amqp.Connection
	Log        zerolog.Logger
}

func InitRbtClient(rbtConfig config.RbtConfig, log zerolog.Logger) (rbtClient *RbtClient, error error) {
	// подключение к rabbitMQ
	connRbtMsg, err := amqp.Dial("amqp://" + rbtConfig.Username + ":" + rbtConfig.Pass + "@" + rbtConfig.Host + ":" + rbtConfig.Port + "/")
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	return &RbtClient{connRbtMsg, log}, nil
}

// метод для открытия канала
func (rbtClient *RbtClient) OpenChannel() (*amqp.Channel, error) {
	ch, err := rbtClient.ConnRbtMsg.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}
