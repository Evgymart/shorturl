package save

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"main/internal/lib/api/response"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(fullUrl string, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const operation = "handlers.url.save.New"
		log = log.With(slog.String("operation", operation))

		decoder := json.NewDecoder(request.Body)
		writer.Header().Set("Content-Type", "application/json")
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			errorMsg := "Failed to decode request body"
			log.Error(errorMsg, err)
			jsonResp, _ := json.Marshal(response.Error(errorMsg))
			_, err := writer.Write(jsonResp)
			if err != nil {
				log.Error("Failed to write a response", err)
			}
			return
		}

		if err = validator.New().Struct(req); err != nil {
			errorMsg := "Failed to write a response"
			log.Error(errorMsg, err)
			jsonResp, _ := json.Marshal(response.Error(errorMsg))

			_, err := writer.Write(jsonResp)
			if err != nil {
				log.Error("Failed to write a response", err)
			}
			return
		}

		alias := req.Alias
		url := req.URL
		err = urlSaver.SaveURL(alias, url)
		if err != nil {
			errorMsg := "Failed to save url"
			log.Error(errorMsg, err)
			jsonResp, _ := json.Marshal(response.Error(errorMsg))
			_, err := writer.Write(jsonResp)
			if err != nil {
				log.Error("Failed to write a response", err)
			}
			return
		}

		jsonResp, _ := json.Marshal(response.OK())
		_, err = writer.Write(jsonResp)
		if err != nil {
			log.Error("Failed to write a response", err)
		}
	}
}
