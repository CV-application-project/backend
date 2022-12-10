package main

import (
	"Backend-Server/timekeeping_service/config"
	"Backend-Server/timekeeping_service/service"
	"Backend-Server/timekeeping_service/store"
	"github.com/jmoiron/sqlx"
)

func newDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newService(cfg *config.Config) (*service.Service, error) {
	db, err := newDB(cfg.MySQL.DSN())
	if err != nil {
		logger.Error(err, "Error connect database")
		return nil, err
	}

	serviceStore := store.New(db)
	return service.NewService(logger, serviceStore, cfg), nil
}
