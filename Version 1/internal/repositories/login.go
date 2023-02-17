package repositories

import (
	"context"
	"fmt"

	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/pkg/postgres"
)

type Login interface {
	Verify(context.Context, string) (entity.Login, error)
}

type loginImplementation struct {
	conn postgres.Adapter
}

func (r *loginImplementation) Verify(ctx context.Context, user string) (login entity.Login, err error) {
	state := true
	query := `UPDATE logins SET logged = $2 WHERE username = $1`

	err = r.conn.QueryRow(ctx, query, user, state).Scan(&login.ID, &login.Username, &login.Role, &login.Logged)

	if err != nil {
		return entity.Login{}, err
	}
	fmt.Println(login)
	return login, nil
}

func NewLoginImplementation(conn postgres.Adapter) *loginImplementation {
	return &loginImplementation{
		conn: conn,
	}
}
