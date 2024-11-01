package handlers

import (
	"encoding/json"
	"errors"
	"github.com/NikCool98/authorization/config"
	"github.com/NikCool98/authorization/storage"
	"net/http"
)

func StartPageHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		m, err := s.GetRandomMotivation()
		if err != nil {
			err := errors.New("Ошибка в получение рандомной фразы")
			config.ErrorResponse.Error = err.Error()
			json.NewEncoder(res).Encode(config.ErrorResponse)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(m); err != nil {
			http.Error(res, `{"error":"Ошибка кодирования JSON"}`, http.StatusInternalServerError)
			return
		}
	}
}
