package handler

import (
	"net/http"
	"techtrain-mission/src/infra/http/request"
	"techtrain-mission/src/infra/http/response"
	"techtrain-mission/src/usecase"

	"github.com/labstack/echo"
)

type UserHandler interface {
	Create() echo.HandlerFunc
	Get() echo.HandlerFunc
	Update() echo.HandlerFunc
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

func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("X-Token")

		if token == "" {
			return c.JSON(http.StatusBadRequest, "token not found")
		}

		targetUser, err := uh.userUsecase.Get(token)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := response.UserGetResponse{
			Name: targetUser.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("X-Token")

		if token == "" {
			return c.JSON(http.StatusBadRequest, "token not found")
		}

		var req request.UserUpdateRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		_, err := uh.userUsecase.Update(req.Name, token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, nil)
	}
}
