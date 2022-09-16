package config

import (
	"DAY_3/controllers"
	"DAY_3/models"
	repositories "DAY_3/repositories"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	c            controllers.Controllers
	onController sync.Once
)

func GetController() controllers.Controllers {
	onController.Do(func() {
		c = controllers.NewControllers(GetRepository())
	})

	return c
}

var (
	repo    repositories.Repositories
	oneRepo sync.Once
)

func GetRepository() repositories.Repositories {
	oneRepo.Do(func() {
		repo = repositories.NewRepositories(GetQuery())
	})

	return repo
}

var (
	qry     *gorm.DB
	oneSync sync.Once
)

func GetQuery() *gorm.DB {

	oneSync.Do(func() {
		dbLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)

		gormConfig := &gorm.Config{
			// enhance performance config
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 dbLogger,
		}

		dsnMaster := os.Getenv("DB_DSN")
		dbMaster, errMaster := gorm.Open(mysql.Open(dsnMaster), gormConfig)
		if errMaster != nil {
			log.Panic(errMaster)
		}

		var err error
		if err = dbMaster.AutoMigrate(
			&models.User{},
			&models.Book{},
		); err != nil {
			log.Fatal(err)
		}

		fmt.Println("success connect to database")
		qry = dbMaster
	})

	return qry
}
