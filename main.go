package main

import (
	"os"

	"sugam-project/api/config"
	"sugam-project/api/controller"
	"sugam-project/api/db/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title mrc-api
// @version 1.0.1
// @description mrc-api
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email info@mrc.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host api.mrc.babulal.com.np
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error in load env file %v", err)
	} else {
		logrus.Info("Successfully loaded env file.")
	}
	conf, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	conn, err := gorm.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		logrus.Fatal("Cannot connect to db ", err)
	}
	pdb := postgres.NewDB(conn)
	server, err := controller.NewServer(pdb.DB)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	port := ":" + os.Getenv("PORT")
	server.Run(port)
}
