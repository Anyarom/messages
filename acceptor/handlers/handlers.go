package handlers

import (
	"acceptor/config"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type WrapperHandler struct {
	Log        zerolog.Logger
	RbtChannel *amqp.Channel
	Cfg        *config.Config
}

func InitWrapperHandler(log zerolog.Logger, cfg *config.Config, rbtChannel *amqp.Channel) *WrapperHandler {
	return &WrapperHandler{Log: log, Cfg: cfg, RbtChannel: rbtChannel}
}
