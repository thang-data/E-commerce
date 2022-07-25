package main

import (
	"github.com/sirupsen/logrus"
	"github.com/thang-data/backend/api/router"
	"github.com/thang-data/backend/config"
	"github.com/thang-data/backend/db"
)

func main() {
	cfg := config.GetConfig()
	db.Init()
	e := router.Init()

	logrus.Fatal(e.Start(":" + cfg.APIPort))
}
