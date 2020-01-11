package main

import (
	"acceptor/config"
	"acceptor/handlers"
	"acceptor/rabbit"
	"github.com/buaazp/fasthttprouter"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {

	// настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Logger()

	// чтение с конфига с помощью библиотеки Viper
	cfg, err := config.InitConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("Ошибка чтения конфига")
	}

	// запуск клиента rabbit
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

	// инициализация структуры WrapperHandler
	wrapperHandler := handlers.InitWrapperHandler(log, cfg, ch)

	// создание роутинга к web-серверу
	router := fasthttprouter.New()
	router.POST("/sms", handlers.InterceptorLogger(wrapperHandler.GetMsgHandler, log))

	// настройка и запуск сервера
	server := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: 20 * 1024 * 1024,
	}
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal().Caller().Err(err).Msg("Ошибка на сервере")
	}
}
