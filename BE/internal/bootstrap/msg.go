// Package bootstrap
package bootstrap

import (
	"BookingApp/BE/internal/consts"
	"BookingApp/BE/pkg/logger"
	"BookingApp/BE/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
