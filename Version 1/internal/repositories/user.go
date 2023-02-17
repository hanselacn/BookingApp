package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/pkg/postgres"
)

type User interface {
	Create(context.Context, *entity.User) (*uuid.UUID, error)
	Verify(context.Context, string) (entity.User, error)
	GrantAdmin(context.Context, string) error
	DemoteToUser(context.Context, string) error
	GetAllUser(context.Context) ([]entity.User, error)
}

type userImplementation struct {
	conn postgres.Adapter
}

func (u userImplementation) Create(ctx context.Context, us *entity.User) (*uuid.UUID, error) {

	query := `
	INSERT INTO users (username, password, name)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	row := u.conn.QueryRow(ctx, query, us.Username, us.Password, us.Name)

	user := entity.User{}

	err := row.Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	querylog := `
	INSERT INTO logins (id, username, password, role, logged)
	VALUES ($1, $2, $3, $4, FALSE)
	`
	_, err = u.conn.Exec(ctx, querylog, &user.ID, us.Username, us.Password, us.Role)

	if err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (r *userImplementation) Verify(ctx context.Context, username string) (user entity.User, err error) {
	query := `SELECT id, username, password, role FROM users WHERE username = $1`

	err = r.conn.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userImplementation) GrantAdmin(ctx context.Context, id string) error {
	fmt.Println(id)
	role := "1"
	query := `UPDATE users SET role = $2 WHERE id = $1`

	rows, err := r.conn.Exec(ctx, query, id, role)

	if err != nil {
		return err
	}
	updated, _ := rows.RowsAffected()
	if updated <= 0 {
		err = fmt.Errorf("user already an admin")
		return err
	}

	return nil
}

func (r *userImplementation) DemoteToUser(ctx context.Context, id string) error {
	role := "0"
	query := `UPDATE users SET role = $2 WHERE id = $1`

	rows, err := r.conn.Exec(ctx, query, id, role)

	if err != nil {
		return err
	}
	updated, _ := rows.RowsAffected()
	if updated <= 0 {
		err = fmt.Errorf("can't demote role user")
		return err
	}

	return nil
}
func (r *userImplementation) GetAllUser(ctx context.Context) ([]entity.User, error) {
	query := `SELECT id, username, password, role, name FROM users`

	queries, err := r.conn.QueryRows(ctx, query)

	if err != nil {
		err = fmt.Errorf("executing query error : %w", err)
		return nil, err
	}

	users := []entity.User{}

	for queries.Next() {
		var user entity.User

		err = queries.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
		)

		if err != nil {
			err = fmt.Errorf("scanning bookings: %w", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

func NewUserImplementation(conn postgres.Adapter) *userImplementation {
	return &userImplementation{
		conn: conn,
	}
}

// query := `
// INSERT INTO users (
// 	username,
// 	password,
// 	name,
// 	role
// )
// VALUES ($1,$2,$3,$4)
// RETURNING id
// `
// fmt.Printf("err1 \n")
// // res, err := u.conn.Exec(ctx, query, us.Username, us.Password, us.Name, us.Role)
// // if err != nil {
// // 	fmt.Printf("err2 \n")
// // 	fmt.Println(err)
// // 	err = fmt.Errorf("create data to db : %w", err)
// // 	return err
// // }
// // fmt.Printf("err3 \n")
// // fmt.Println(res)

// row := u.conn.QueryRow(ctx, query, us.Username, us.Password, us.Name, us.Role)

// user := entity.User{}

// err = row.Scan(&user.ID)

// if err != nil{

// }

// return &user.ID, nil
// user.ID, err = res.LastInsertId()
// if err != nil {
// 	fmt.Printf("err4 \n")
// 	err = fmt.Errorf("inserting last id to db : %w", err)
// 	return err
// }
// fmt.Printf("err5 \n")
// return nil
// return &user.ID, nil
