package gateway

import (
	"github.com/go-pg/pg/v9"
	"github.com/kallydev/yogurt/common/config"
	"github.com/kallydev/yogurt/common/database"
	_ "github.com/lib/pq"
	"log"
)

var (
	Conf *config.Config
	DB   *pg.DB
)

const confPath = "config/config_service-gateway.yaml"

func init() {
	var err error
	if Conf, err = config.ParseConfigFile(confPath); err != nil {
		log.Fatalln(err)
	} else {
		DB = database.DialPostgres(
			Conf.Postgres.Username, Conf.Postgres.Password,
			Conf.Postgres.Host, Conf.Postgres.Port, Conf.Postgres.Database, nil,
		)
	}
}
