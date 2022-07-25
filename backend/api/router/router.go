package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thang-data/backend/api/api_pkg"
	"github.com/thang-data/backend/api/handler"
	"github.com/thang-data/backend/config"
	"github.com/thang-data/backend/log/echologrus"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.Logger = echologrus.Logger()
	e.Use(api_pkg.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//e.GET("/", handler.Health)
	//e.GET("/health", handler.Health)
	//e.GET("/hello", handler.Hello)
	//e.GET("/hello-session", handler.Hello, api_pkg.Session)
	if config.GetConfig().ApiMode {
		e.POST("/admin/signup-by-email-password", handler.SignupAdminByEmailPassword)
		e.POST("/admin/signup-information", handler.SignupInformation)
		e.POST("/admin/login-by-email-password", handler.LoginAdminByEmailPassword)
	}

	return e
}
func LoggerSkipper(c echo.Context) bool {
	if c.Path() == "/health" {
		return true
	}

	return false
}
