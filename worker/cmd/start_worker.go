package cmd

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"worker/config"
	"worker/postgres"
	"worker/rabbit"
)

var startWorker = &cobra.Command{
	Use:   "start_worker",
	Short: "start worker",
	Run: func(cmd *cobra.Command, args []string) {
		runWorker()
	},
}

func runWorker() {
	// зададим настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Logger()

	// чтение с конфига с помощью библиотеки Viper
	cfg, err := config.InitConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("Ошибка")
	}

	// создание клиента для Pg
	pgxClient, err := postgres.InitPgxClient(cfg.PgConfig, log)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("клиент pgxClient не инициализирован")
	}

	// создание клиента для rabbit
	rbtClient, err := rabbit.InitRbtClient(cfg.RbtConfig, log)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("клиент rbtClient не инициализирован")
	}
	defer func() {
		if err := rbtClient.Close(); err != nil {
		}
	}()

	// откроем канал
	ch, err := rbtClient.OpenChannel()
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	msgs, err := rabbit.ConsumerMsg(ch, *pgxClient)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}

	// создание структуры для парсинга боди с запроса
	type Message struct {
		Phone string `json:"phone"`
		Text  string `json:"text"`
	}

	// читаем сообщения из очереди и пишем в бд
	for d := range msgs {
		// парсинг сообщения
		var message Message
		err := json.Unmarshal(d.Body, &message)
		if err != nil {
			log.Debug().Caller().Err(err).Msg("")
			err := d.Nack(false, true)
			if err != nil {
				log.Debug().Caller().Err(err).Msg("")
				continue
			}
		}

		// записываем сообщение в pg
		err = pgxClient.DbInsertMessage(message.Phone, message.Text)
		if err != nil {
			log.Debug().Caller().Err(err).Msg("")
			err := d.Nack(false, true)
			if err != nil {
				log.Debug().Caller().Err(err).Msg("")
				continue
			}
		}

		// подтвердим, что сообщение обработано
		err = d.Ack(false)
		if err != nil {
			log.Debug().Caller().Err(err).Msg("")
		}
	}
}
