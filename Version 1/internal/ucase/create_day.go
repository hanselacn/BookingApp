package ucase

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type createDay struct {
	dayRepository repositories.Day
}

func (cd createDay) Serve(data *appctx.Data) appctx.Response {
	a := &entity.Day{}

	err := data.Cast(&a)

	if err != nil {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	user := entity.User{}

	user, err = cd.dayRepository.VerifyDay(data.Request.Context(), a.CreatedBy)

	if user.Username != a.CreatedBy {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	id, err := cd.dayRepository.NewCreateDay(data.Request.Context(), a)

	if err != nil {
		fmt.Println(err)
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	type resp struct {
		ID uuid.UUID `json:"id"`
	}

	var res resp
	res.ID = *id

	return *appctx.NewResponse().WithCode(http.StatusCreated).WithStatus("SUCCESS").WithEntity("Users").WithState("AddingUserSuccess").WithMessage("Adding User Success").WithData(res)
}

func NewCreateDay(dayRepository repositories.Day) *createDay {
	return &createDay{
		dayRepository: dayRepository,
	}
}
