package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/thang-data/backend/usecase"
	"net/http"
)

func SignupAdminByEmailPassword(c echo.Context) error {
	email := c.FormValue("email")
	lastName := c.FormValue("lastName")
	firstName := c.FormValue("firstName")
	profile, err := usecase.SignupAdminByEmailPassword(email, lastName, firstName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func SignupInformation(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	repeatPassword := c.FormValue("repeatPassword")

	admin, err := usecase.SignupInformation(email, password, repeatPassword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, admin)
}

func LoginAdminByEmailPassword(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	session, err := usecase.LoginAdminByEmailPassword(email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, session)
}
