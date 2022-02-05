package handler

import (
	"encoding/json"
	"net/http"

	myError "go-gacha-server/src/core/error"
	"go-gacha-server/src/presen/request"
	"go-gacha-server/src/presen/response"
	"go-gacha-server/src/usecase"

	"go.uber.org/zap"
)

type UserHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	var req request.UserCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		zap.Error(err)
	}

	createdUser, err := uh.userUsecase.Create(r.Context(), req.Name)
	if err != nil {
		zap.Error(err)
	}

	res := response.UserCreateResponse{
		Token: createdUser.Token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
	}
}

func (uh *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	token := r.Header.Get("X-Token")
	if token == "" {
		zap.Error(myError.ErrTokenNotFound)
	}

	targetUser, err := uh.userUsecase.Get(r.Context(), token)

	if err != nil {
		zap.Error(err)
	}

	res := response.UserGetResponse{
		Name: targetUser.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
	}
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(myError.ErrMethodNotFound)
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		zap.Error(myError.ErrTokenNotFound)
	}

	var req request.UserUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		zap.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	_, err = uh.userUsecase.Update(r.Context(), req.Name, token)
	if err != nil {
		zap.Error(err)
	}
}
