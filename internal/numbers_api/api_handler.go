package numbers_api

import (
	"context"
	"github.com/jeroldleslie/go-coding-challenge-trivia/internal/utils"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type APIHandler struct {
	API APIImpl
}

func NewAPIHandler(impl APIImpl) *APIHandler {
	return &APIHandler{impl}
}

func (a *APIHandler) Ping(c echo.Context) error {
	response := utils.Response{
		StatusCode: http.StatusOK,
		Data:       http.StatusText(http.StatusOK),
	}
	return response.Send(&c)
}

func (a *APIHandler) GetDetails(c echo.Context) error {
	ctx := a.clientContextAndAllowCors(c)

	number := c.QueryParam("number")
	text := c.QueryParam("text")

	numbers, err := a.API.GetDetails(ctx, number, text)
	return respond(&c, numbers, err)
}

func respond(c *echo.Context, data interface{}, err error) error {
	logrus.Tracef("err: %+v", err)
	response := utils.Response{
		StatusCode: errorToHTTPStatusCode(err),
		Data:       data,
		Error:      err,
	}
	return response.Send(c)
}

func (a *APIHandler) clientContextAndAllowCors(c echo.Context) context.Context {
	utils.AllowCors(c)
	const clientContextTimeout = 40 * time.Second
	return utils.GenerateTimeoutContext(c, clientContextTimeout)
}

func errorToHTTPStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case ErrBadInput:
		return http.StatusBadRequest
	case ErrNoDataAvailable:
		return http.StatusNoContent
	}
	return http.StatusInternalServerError
}
