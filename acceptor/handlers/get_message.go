package handlers

import (
	"acceptor/rabbit"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// структура для парсинга запроса
type AcceptorMsg struct {
	Method  string `json:"method"`
	Address string `json:"address"`
	Body    string `json:"body"`
}

func (wrapHandler *WrapperHandler) GetMsgHandler(ctx *fasthttp.RequestCtx) {

	// получение боди из запроса
	reqBody := ctx.Request.Body()

	// провалидируем что структура запроса верная
	var acceptorMsg AcceptorMsg
	err := json.Unmarshal(reqBody, &acceptorMsg)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Err(err).Msg("")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// reqBody отправим в rabbit
	err = rabbit.SenderToWorker(wrapHandler.RbtChannel, reqBody)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Err(err).Msg("")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	}
}
