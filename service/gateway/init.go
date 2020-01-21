package gateway

import (
	"database/sql"
	"github.com/kallydev/yogurt/common/config"
	"github.com/kallydev/yogurt/common/database"
	_ "github.com/lib/pq"
	"log"
)

var (
	Conf *config.Config
	DB   *sql.DB
)

const confPath = "config/config_service-gateway.yaml"

func init() {
	var err error
	if Conf, err = config.ParseConfigFile(confPath); err != nil {
		log.Fatalln(err)
	} else if DB, err = database.DialPostgres(
		Conf.Postgres.Username, Conf.Postgres.Password,
		Conf.Postgres.Host, Conf.Postgres.Port, Conf.Postgres.Database, nil,
	); err != nil {
		log.Fatalln(err)
	}
}
