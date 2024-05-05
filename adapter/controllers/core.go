package controllers

import (
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// responses
type (
	ErrorResponse struct {
		Message string `json:"message"`
	}

	EmptyResponse struct{}
)

var (
	DefaultResponse = EmptyResponse{}
)

// errors
func ErrorHandle(c echo.Context, err error) error {
	switch err.(type) {
	case *errors.NotFoundError:
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	case *errors.AlreadyExistsError:
		return c.JSON(http.StatusConflict, ErrorResponse{Message: err.Error()})
	case *errors.AuthenticationError:
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
	case *errors.AuthorizationError:
		return c.JSON(http.StatusForbidden, ErrorResponse{Message: err.Error()})
	case *errors.SystemError:
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("Internal server error. message:%s", err.Error())})
}

// validaton
// type CustomValidator struct {
// 	validator *validator.Validate
// }

// // SEE:https://github.com/go-playground/validator
// // https://pkg.go.dev/gopkg.in/go-playground/validator.v9#section-readme
// // https://qiita.com/RunEagler/items/ad79fc860c3689797ccc
// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }

// func NewValidator() echo.Validator {
// 	validator := validator.New()
// 	validator.RegisterValidation("date", isDate)

// 	return &CustomValidator{validator: validator}
// }

// func isDate(fl validator.FieldLevel) bool {
// 	_, err := time.Parse("2006-01-02", fl.Field().String())
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// requests
type (
	RequestMapper struct{}
)

func (it *RequestMapper) Parse(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	// if err := c.Validate(i); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	// }
	return nil
}

func (it *RequestMapper) GetSessionKey(c echo.Context) (models.SessionKey, error) {
	sk := c.Request().Header.Get("x-session-key")

	if sk == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "x-session-key is required."})
	}

	return models.SessionKey(sk), nil
}

func (it *RequestMapper) GetMasterKey(c echo.Context) (models.MasterKey, error) {
	mk := c.Request().Header.Get("x-master-key")

	if mk == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "x-master-key is required."})
	}

	return models.MasterKey(mk), nil
}

func NewRequestMapper() RequestMapper {
	return RequestMapper{}
}

// converter
func ToDate(src string) time.Time {
	dt, _ := time.Parse("2006-01-02", src)
	return dt
}

func ToString(t time.Time) string {
	return t.Format("2006-01-02")
}
