package config

import (
	"Backend-Server/library/database"
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	database.Base   `mapstructure:",squash"`
	MySQL           database.DBConfig `json:"mysql" mapstructure:"mysql"`
	MigrationFolder string            `json:"migration_folder" mapstructure:"migration_folder"`
}

func loadDefaultConfig() *Config {
	return &Config{
		Base:            *database.DefaultBaseConfig(),
		MySQL:           database.MySQLDefaultConfig(),
		MigrationFolder: "file://user_service/sql/migrations",
	}
}

func Load() (*Config, error) {
	c := loadDefaultConfig()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./cv_service")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Read config file failed. ", err)

		configBuffer, err := json.Marshal(c)

		if err != nil {
			return nil, err
		}

		if err = viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
			return nil, err
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
