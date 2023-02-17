// Package bootstrap
package bootstrap

import (
	"BookingApp/BE/internal/appctx"
	"BookingApp/BE/pkg/logger"
	"BookingApp/BE/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}
