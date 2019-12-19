package handlers

import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

// создание middleware
func InterceptorLogger(myAnyHandler fasthttp.RequestHandler, log zerolog.Logger) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// логирование запроса
		log.Debug().Bytes("request body", ctx.Request.Body()).Msg("")

		// передаем управление нашему основному Handler
		myAnyHandler(ctx)

		// логирование ответа
		//log.Debug().Bytes("resp", ctx.Response.Body()).Msg("")

	}
}
