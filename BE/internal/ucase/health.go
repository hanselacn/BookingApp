package ucase

import (
	"BookingApp/BE/internal/appctx"
	"BookingApp/BE/internal/consts"
	"BookingApp/BE/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("ok")
}
