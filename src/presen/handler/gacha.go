package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	myError "techtrain-mission/src/core/error"
	"techtrain-mission/src/presen/request"
	"techtrain-mission/src/presen/response"
	"techtrain-mission/src/usecase"

	"go.uber.org/zap"
)

type GachaHandler interface {
	Draw(http.ResponseWriter, *http.Request)
}

type gachaHandler struct {
	gachaUsecase usecase.GachaUsecase
}

func NewGachaHandler(gachaUsecase usecase.GachaUsecase) GachaHandler {
	return &gachaHandler{gachaUsecase: gachaUsecase}
}

func (gh *gachaHandler) Draw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	var req request.GachaDrawRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		zap.Error(err)
	}

	characters, err := gh.gachaUsecase.Draw(r.Context(), req.Times)
	if err != nil {
		zap.Error(err)
	}

	var results []response.GachaResult
	for _, ce := range characters {
		result := response.GachaResult{
			CharacterID: strconv.Itoa(ce.Id),
			Name:        ce.Name,
		}
		results = append(results, result)
	}

	res := response.GachaDrawResponse{
		Results: results,
	}

	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
	}
}
