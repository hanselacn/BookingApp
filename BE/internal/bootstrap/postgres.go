// Package bootstrap
package bootstrap

import (
	"time"

	config "BookingApp/BE/internal/appctx"

	"BookingApp/BE/pkg/logger"
	"BookingApp/BE/pkg/postgres"
)

func RegistryPostgreSQL(cfg *config.Database, timezone string) postgres.Adapter {
	db, err := postgres.NewPostgreSQL(&postgres.Config{
		Host:         cfg.Host,
		Name:         cfg.Name,
		Password:     cfg.Pass,
		Port:         cfg.Port,
		User:         cfg.User,
		Timeout:      time.Duration(cfg.TimeoutSecond) * time.Second,
		MaxOpenConns: cfg.MaxOpen,
		MaxIdleConns: cfg.MaxIdle,
		MaxLifetime:  time.Duration(cfg.MaxLifeTimeMS) * time.Millisecond,
		Charset:      cfg.Charset,
		TimeZone:     timezone,
	})

	if err != nil {
		logger.Fatal(
			err,
			logger.EventName("db"),
			logger.Any("host", cfg.Host),
			logger.Any("port", cfg.Port),
		)
	}

	return db
}

func RegistryPostgreSQLMasterSlave(cfgWrite *config.Database, cfgRead *config.Database, timezone string) postgres.Adapter {
	db, err := postgres.NewPostgresMasterSlave(
		&postgres.Config{
			Host:         cfgWrite.Host,
			Name:         cfgWrite.Name,
			Password:     cfgWrite.Pass,
			Port:         cfgWrite.Port,
			User:         cfgWrite.User,
			Timeout:      time.Duration(cfgWrite.TimeoutSecond) * time.Second,
			MaxOpenConns: cfgWrite.MaxOpen,
			MaxIdleConns: cfgWrite.MaxIdle,
			MaxLifetime:  time.Duration(cfgWrite.MaxLifeTimeMS) * time.Millisecond,
			Charset:      cfgWrite.Charset,
			TimeZone:     timezone,
		},

		&postgres.Config{
			Host:         cfgRead.Host,
			Name:         cfgRead.Name,
			Password:     cfgRead.Pass,
			Port:         cfgRead.Port,
			User:         cfgRead.User,
			Timeout:      time.Duration(cfgRead.TimeoutSecond) * time.Second,
			MaxOpenConns: cfgRead.MaxOpen,
			MaxIdleConns: cfgRead.MaxIdle,
			MaxLifetime:  time.Duration(cfgRead.MaxLifeTimeMS) * time.Millisecond,
			Charset:      cfgRead.Charset,
			TimeZone:     timezone,
		},
	)

	if err != nil {
		logger.Fatal(err,
			logger.EventName("db"),
			logger.Any("host_read", cfgRead.Host),
			logger.Any("port_read", cfgRead.Port),
			logger.Any("host_write", cfgWrite.Host),
			logger.Any("port_write", cfgWrite.Port),
		)
	}

	return db
}
