package campaign

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Settings struct {
	Port          string `envconfig:"APP_PORT" default:"9000"`
	DatabaseHost  string `envconfig:"DATABASE_HOST" default:"elasticsearch"`
	DatabasePort  int    `envconfig:"DATABASE_PORT" default:"9200"`
	DatabaseProto string `envconfig:"DATABASE_PROTO" default:"http"`
	DatabaseIndex string `envconfig:"DATABASE_INDEX_FORMAT" default:"clicks"`
	MappingPath   string `envconfig:"DATABASE_MAPPING_PATH" default:"/app/bin/click.json"`
}

func NewSettings() *Settings {
	Config := &Settings{}
	if err := envconfig.Process("", Config); err != nil {
		log.Fatal(err)
	}
	return Config
}
