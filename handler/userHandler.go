package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/takuya911/mcrsvc_user/conf"
	"github.com/takuya911/mcrsvc_user/domain"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	usecase      domain.UserUsecase
	passwordHash int
}

// NewUserHandler function
func NewUserHandler(e *echo.Echo, u domain.UserUsecase) {
	handler := &userHandler{
		usecase: u,
	}
	e.GET("/:user_id", handler.GetUserByID)
	e.POST("/user", handler.StoreUser)
	e.POST("/:user_id", handler.UpdateUser)
	e.DELETE("/:user_id", handler.DeleteUser)
}

func (h *userHandler) GetUserByID(e echo.Context) error {
	userID, _ := strconv.Atoi(e.Param("user_id"))
	if userID < 1 {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}

	etx := e.Request().Context()
	user, err := h.usecase.GetByID(etx, userID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (h *userHandler) StoreUser(e echo.Context) (err error) {
	etx := e.Request().Context()

	var user domain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isFormValid(&user); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	hash, err := passwordHash(user.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = hash

	err = h.usecase.Store(etx, &user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(e echo.Context) error {
	userID, err := strconv.Atoi(e.Param("user_id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}
	etx := e.Request().Context()

	var user domain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	user.ID = userID

	var ok bool
	if ok, err = isFormValid(&user); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	hash, err := passwordHash(user.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = hash

	err = h.usecase.Update(etx, &user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

func (h *userHandler) DeleteUser(e echo.Context) error {
	userID, err := strconv.Atoi(e.Param("user_id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}

	etx := e.Request().Context()
	if err := h.usecase.Delete(etx, userID); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.NoContent(http.StatusNoContent)
}

func isFormValid(u *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

// パスワードハッシュを作る
func passwordHash(pw string) (string, error) {
	hashNum, err := strconv.Atoi(os.Getenv("HASH_NUMBER"))
	if err != nil {
		return "", err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pw), hashNum)
	if err != nil {
		return "", err
	}
	return string(hashedPass), err
}
