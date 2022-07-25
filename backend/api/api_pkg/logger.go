package api_pkg

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}
			p := req.URL.Path
			bytesIn := req.Header.Get(echo.HeaderContentLength)
			if bytesIn == "" {
				bytesIn = "0"
			}
			fields := logrus.Fields{
				"bytes_in":   bytesIn,
				"bytes_out":  strconv.FormatInt(res.Size, 10),
				"host":       req.Host,
				"latency":    strconv.FormatInt(stop.Sub(start).Nanoseconds(), 10),
				"method":     req.Method,
				"path":       p,
				"remote_ip":  c.RealIP(),
				"status":     res.Status,
				"uri":        req.RequestURI,
				"user_agent": req.UserAgent(),
			}
			if err != nil {
				logrus.WithFields(fields).Info(err)
			} else {
				logrus.WithFields(fields).Info()
			}

			return err
		}
	}
}
