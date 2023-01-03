package registry

import (
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OrmService() *di.Def {
	return &di.Def{
		Name: OrmServiceDefinition,
		Build: func(ctn di.Container) (interface{}, error) {
			dsn := "----"

			log.Debug("connection to database ...")

			gormConfig := &gorm.Config{}

			db, err := gorm.Open(mysql.Open(dsn), gormConfig)

			if err != nil {
				log.Infof("could not connect to database: %s", err.Error())
				return nil, err
			}

			log.Info("successfully connected to database")

			return db, nil
		},
	}
}
