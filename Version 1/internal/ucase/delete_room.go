package ucase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type deleteRoom struct {
	repo repositories.Room
}

func NewDeleteRoom(repo repositories.Room) *deleteRoom {
	return &deleteRoom{repo: repo}
}

func (u *deleteRoom) Serve(data *appctx.Data) appctx.Response {
	id := mux.Vars(data.Request)["roomID"]
	err := u.repo.DeleteRoom(data.Request.Context(), id)

	if err != nil {
		err := fmt.Errorf("deleting room: %w", err)
		fmt.Println(err)

		return *appctx.NewResponse().WithCode(http.StatusNotFound).WithStatus("FAIL").WithEntity("Room").WithState("DeleteRoomFail").WithMessage("Delete Room Fail").WithError(err.Error())
	}

	return *appctx.NewResponse().WithCode(http.StatusOK).WithStatus("SUCCESS").WithEntity("Room").WithState("DeleteRoomSuccess").WithMessage("Delete Room Success")
}
