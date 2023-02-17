// Package ucase
package ucase

import (
	"testing"

	"BookingApp/BE/internal/appctx"
	"BookingApp/BE/internal/consts"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Serve(t *testing.T) {
	svc := NewHealthCheck()

	t.Run("test health check", func(t *testing.T) {
		result := svc.Serve(&appctx.Data{})

		assert.Equal(t, appctx.Response{
			Code:    consts.CodeSuccess,
			Message: "ok",
		}, result)
	})
}
