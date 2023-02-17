package ucase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type updateAdmin struct {
	repo repositories.User
}

func NewUpdateAdmin(repo repositories.User) *updateAdmin {
	return &updateAdmin{repo: repo}
}

func (u *updateAdmin) Serve(data *appctx.Data) appctx.Response {
	id := mux.Vars(data.Request)["userID"]
	err := u.repo.GrantAdmin(data.Request.Context(), id)

	if err != nil {
		err := fmt.Errorf("grant to admin: %w", err)
		fmt.Println(err)

		return *appctx.NewResponse().WithCode(http.StatusNotFound).WithStatus("FAIL").WithEntity("User").WithState("updateAdminFail").WithMessage("Update Admin Fail").WithError(err.Error())
	}

	return *appctx.NewResponse().WithCode(http.StatusOK).WithStatus("SUCCESS").WithEntity("User").WithState("updateAdminSuccess").WithMessage("Update Admin Success")
}
