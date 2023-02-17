package ucase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type updateUser struct {
	repo repositories.User
}

func NewUpdateUser(repo repositories.User) *updateUser {
	return &updateUser{repo: repo}
}

func (u *updateUser) Serve(data *appctx.Data) appctx.Response {
	id := mux.Vars(data.Request)["userID"]
	err := u.repo.DemoteToUser(data.Request.Context(), id)

	if err != nil {
		err := fmt.Errorf("demote to user: %w", err)
		fmt.Println(err)

		return *appctx.NewResponse().WithCode(http.StatusNotFound).WithStatus("FAIL").WithEntity("User").WithState("updateUserFail").WithMessage("Update User Fail").WithError(err.Error())
	}

	return *appctx.NewResponse().WithCode(http.StatusOK).WithStatus("SUCCESS").WithEntity("User").WithState("updateUserSuccess").WithMessage("Update User Success")
}
