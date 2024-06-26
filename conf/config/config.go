package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type MysqlConfig struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
}

type ServerConfig struct {
	ServerPort int
	MasterKey  string
}

var Mysql MysqlConfig
var Server ServerConfig

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Mysql = MysqlConfig{
		DbDriverName:   cfg.Section("db").Key("db_driver_name").String(),
		DbName:         cfg.Section("db").Key("db_name").String(),
		DbUserName:     cfg.Section("db").Key("db_user_name").String(),
		DbUserPassword: cfg.Section("db").Key("db_user_password").String(),
		DbHost:         cfg.Section("db").Key("db_host").String(),
		DbPort:         cfg.Section("db").Key("db_port").String(),
	}

	Server = ServerConfig{
		ServerPort: cfg.Section("server").Key("server_port").MustInt(),
		MasterKey:  cfg.Section("server").Key("master_key").String(),
	}
}
