package utils

import (
	"context"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func AllowCors(c echo.Context) {
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
}

func GenerateTimeoutContext(c echo.Context, timeout time.Duration) context.Context {
	if timeout == 0 {
		//logrus.Errorf("missing timeout", errors.Errorf("for path %+v", c.Request().URL.Path))
		log.Printf("missing timeout %+v", errors.Errorf("for path %+v", c.Request().URL.Path))
		return c.Request().Context()
	}
	logrus.Infof("timeout context for path %s: %s", c.Request().URL.Path, timeout)
	ctx, _ := context.WithTimeout(c.Request().Context(), timeout)
	ctx = context.WithValue(ctx, "QueryParams", c.QueryParams())
	return ctx
}
