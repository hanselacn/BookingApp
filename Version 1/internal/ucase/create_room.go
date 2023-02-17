package ucase

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type createRoom struct {
	roomRepository repositories.Room
}

func (cr createRoom) Serve(data *appctx.Data) appctx.Response {
	a := &entity.Room{}

	err := data.Cast(&a)

	if err != nil {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	user := entity.User{}

	user, err = cr.roomRepository.VerifyRoom(data.Request.Context(), a.CreatedBy)

	if user.Username != a.CreatedBy {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	room := entity.Room{}

	room, _ = cr.roomRepository.VerifyRoomEntity(data.Request.Context(), a.RoomName)

	if room.RoomName == a.RoomName {
		err = fmt.Errorf("room already exist")
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	id, err := cr.roomRepository.Create(data.Request.Context(), a)

	if err != nil {
		fmt.Println(err)
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Room").WithState("AddRoomFailed").WithMessage("Adding Room Failed").WithError(err.Error())
	}

	type resp struct {
		ID uuid.UUID `json:"id"`
	}

	var res resp
	res.ID = *id

	return *appctx.NewResponse().WithCode(http.StatusCreated).WithStatus("SUCCESS").WithEntity("Room").WithState("AddingRoomSuccess").WithMessage("Adding Room Success").WithData(res)
}

func NewCreateRoom(roomRepository repositories.Room) *createRoom {
	return &createRoom{
		roomRepository: roomRepository,
	}
}
