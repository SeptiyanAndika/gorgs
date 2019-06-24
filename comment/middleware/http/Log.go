package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func RoutingLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		start := time.Now()
		req := c.Request()
		res := c.Response()
		c.Set("RequestURI", req.RequestURI)

		logrus.WithFields(logrus.Fields{
			"requestId":   res.Header().Get(echo.HeaderXRequestID),
			"method":      req.Method,
			"timeRequest": start.Format("Mon Jan _2 15:04:05 2006"),
		}).Info(req.RequestURI)
		c.Set("requestURI", req.RequestURI)

		if err = next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()
		status := res.Status

		logrus.WithFields(logrus.Fields{
			"requestId":    res.Header().Get(echo.HeaderXRequestID),
			"responseTime": stop.Sub(start).String(),
			"status":       status,
		}).Info(req.RequestURI)

		return
	}
}
