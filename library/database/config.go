package database

import (
	"fmt"
	"net/url"
)

type DBConfig interface {
	String() string
	DSN() string
}

// Config used to set base config for all database types.
type Config struct {
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Database string `json:"database" mapstructure:"database" yaml:"database"`
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`
	Username string `json:"username" mapstructure:"username" yaml:"username"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
	Options  string `json:"options" mapstructure:"options" yaml:"options"`
}

func (c Config) DSN() string {
	options := c.Options
	if options != "" {
		if options[0] != '?' {
			options = "?" + options
		}
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database,
		options)
}

type MySQLConfig struct {
	Config `mapstructure:",squash"`
}

func (c MySQLConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DSN())
}

func MySQLDefaultConfig() MySQLConfig {
	return MySQLConfig{Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "smart_work_time",
		Username: "root",
		Password: "mysqlPassword",
		Options:  "?parseTime=true",
	}}
}
