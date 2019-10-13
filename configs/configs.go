package configs

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name             string `mapstructure:"NAME"`
	Listen           int    `mapstructure:"LISTEN"`
	DBType           string `mapstructure:"DB_TYPE"`
	DBConnectString  string `mapstructure:"DB_CONNECT_STRING"`
	RDBConnectString string `mapstructure:"RDB_CONNECT_STRING"`
	// ElasticSearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
}

func GetConf() (*ServerConfig, error) {
	var conf ServerConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.New("invalid values of viper")
	}
	return &conf, nil
}
