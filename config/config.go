package config

import (
	"go-practice/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").MustString("mysql"),
		DbName:    cfg.Section("db").Key("name").MustString("local_db"),
		LogFile:   cfg.Section("web").Key("logfile").MustString("webapp.log"),
		Static:    cfg.Section("web").Key("static").MustString("static"),
	}
}
