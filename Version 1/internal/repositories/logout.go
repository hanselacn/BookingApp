package repositories

import (
	"context"
	"fmt"

	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/pkg/postgres"
)

type Logout interface {
	Verify(context.Context, string) (entity.Login, error)
}

type logoutImplementation struct {
	conn postgres.Adapter
}

func (r *logoutImplementation) Verify(ctx context.Context, user string) (login entity.Login, err error) {
	state := false
	query := `UPDATE logins SET logged = $2 WHERE username = $1`

	err = r.conn.QueryRow(ctx, query, user, state).Scan(&login.ID, &login.Username, &login.Password, &login.Role, &login.Logged)

	fmt.Println(err)

	if err != nil {
		return entity.Login{}, err
	}
	fmt.Println(login)
	return login, nil
}

func NewLogoutImplementation(conn postgres.Adapter) *logoutImplementation {
	return &logoutImplementation{
		conn: conn,
	}
}
