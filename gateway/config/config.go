package config

import (
	"Backend-Server/library/database"
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type ClientHost struct {
	UserService        string `json:"user_service" mapstructure:"user_service"`
	CVService          string `json:"cv_service" mapstructure:"cv_service"`
	TimekeepingService string `json:"timekeeping_service" mapstructure:"timekeeping_service"`
}

type Config struct {
	database.Base   `mapstructure:",squash"`
	MySQL           database.DBConfig `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	ClientHost      ClientHost        `json:"client_host" mapstructure:"client_host"`
	MigrationFolder string            `json:"migration_folder" yaml:"migration_folder" mapstructure:"migration_folder"`
}

func loadDefaultConfig() *Config {
	return &Config{
		Base:            *database.DefaultBaseConfig(),
		MySQL:           database.MySQLDefaultConfig(),
		ClientHost:      ClientHost{},
		MigrationFolder: "file://gateway/sql/migrations",
	}
}

func Load() (*Config, error) {
	c := loadDefaultConfig()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./gateway")
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
