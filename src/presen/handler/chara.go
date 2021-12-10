package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	myError "techtrain-mission/src/core/error"
	"techtrain-mission/src/presen/response"
	"techtrain-mission/src/usecase"

	"go.uber.org/zap"
)

type CharaHandler interface {
	List(http.ResponseWriter, *http.Request)
}

type charaHandler struct {
	charaUsecase usecase.CharaUsecase
}

func NewCharaHandler(charaUsecase usecase.CharaUsecase) CharaHandler {
	return &charaHandler{charaUsecase: charaUsecase}
}

func (ch *charaHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		log.Println("token not found")
	}

	ucEntities, err := ch.charaUsecase.List(token)

	if err != nil {
		zap.Error(err)
	}

	var characters []response.UserCharacter
	for _, uce := range ucEntities {
		character := response.UserCharacter{
			UserCharacterID: strconv.Itoa(uce.User.Id),
			CharacterID:     strconv.Itoa(uce.Chara.Id),
			Name:            uce.Chara.Name,
		}
		characters = append(characters, character)
	}

	res := response.CharacterListResponse{
		Characters: characters,
	}

	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
	}
}
