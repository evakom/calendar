/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 27.10.2019 12:32
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"flag"
	"github.com/evakom/calendar/internal/configs"
	"github.com/evakom/calendar/internal/dbs"
	"github.com/evakom/calendar/internal/domain/calendar"
	"log"
	"os"
)

// Constants
const (
	EnvCalendarConfigPath = "CALENDAR_CONFIG_PATH"
)

func main() {

	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	confPath := os.Getenv(EnvCalendarConfigPath)
	if confPath == "" {
		confPath = *configFile
	}

	conf := configs.NewConfig(confPath)
	if err := conf.ReadParameters(); err != nil {
		log.Fatalln(err)
	}

	db, err := dbs.NewDB(conf.DBType)
	if db == nil {
		log.Fatalf("unsupported DB type: %s\n", conf.DBType)
	}

	if err != nil {
		log.Fatalf("Open DB: %s, error: %s \n", conf.DBType, err)
	}

	calendar.PrintTestData(db)
}
