package ucase

import (
	"fmt"
	"net/http"

	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/consts"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/internal/helper/encrypt"
	"gitlab.privy.id/go_graphql/internal/repositories"
	"gitlab.privy.id/go_graphql/pkg/logger"
)

type login struct {
	loginRepository repositories.Login
}

func NewLogin(loginRepository repositories.Login) *login {
	return &login{
		loginRepository: loginRepository,
	}
}
func (u *login) Serve(data *appctx.Data) appctx.Response {
	p := &entity.Login{}

	err := data.Cast(&p)

	if err != nil {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("User").WithState("UserAddFailed").WithMessage("Adding User Fail").WithError(err.Error())
	}

	p.Password, err = encrypt.HashPassword(p.Password)

	if err != nil {
		logger.Error(err)
		return *appctx.NewResponse().WithCode(consts.CodeBadRequest).WithStatus("FAIL").WithEntity("User").WithState("UserAddFailed").WithMessage("Adding User Fail").WithError(err.Error())
	}

	user, err := u.loginRepository.Verify(data.Request.Context(), p.Username)

	success := encrypt.CheckPasswordHash(user.Password, p.Password)

	if !success {
		return *appctx.NewResponse().WithCode(http.StatusUnauthorized).WithStatus("FAIL").WithEntity("Authentications").WithState("AuthenticationFail").WithMessage("Authentication Fail")
	}

	if err != nil {
		fmt.Println(err)
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Users").WithState("UserAddFailed").WithMessage("Adding User Fail").WithError(err.Error())
	}
	return *appctx.NewResponse().WithCode(http.StatusOK).WithStatus("SUCCESS").WithEntity("Users").WithState("AddingUserSuccess").WithMessage("Adding User Success").WithData(user)
}
