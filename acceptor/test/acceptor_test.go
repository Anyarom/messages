package test

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"testing"
)

func TestGetMessages(t *testing.T) {

	// зададим настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Logger()

	// создадание httpClient
	client := http.DefaultClient

	// боди в запросе
	bodyRequest := []byte(`{
                   "phone": "1234",
                   "text": "hello"
                   }`)

	body := bytes.NewReader(bodyRequest)

	// отправим запрос
	resp, err := client.Post("http://127.0.0.1:8080/sms", "application/json", body)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		t.Fail()
	}

	// проверим статус ответа
	if resp.StatusCode != fasthttp.StatusOK {
		t.Fail()
	}
}
