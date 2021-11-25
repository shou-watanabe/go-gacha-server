package handler

import (
	"net/http"
	_ "techtrain-mission/src/domain/entity"
	"techtrain-mission/src/infra/http/request"
	"techtrain-mission/src/usecase"

	"github.com/labstack/echo"
)

type UserHandler interface {
	Create() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type responseTask struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (uh *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UserCreateRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := uh.userUsecase.Create(req.Name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:    createdUser.Id,
			Name:  createdUser.Name,
			Token: createdUser.Token,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
