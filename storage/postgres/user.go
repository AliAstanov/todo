package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repoi.UserRepoI {
	return &UserRepo{
		conn: conn,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error {
	query := `
		INSET INTO users(
		user_id,
		fullname, 
		username, 
		gmail 
	)VALUES (
		$1 ,$2, $3, $4, $5
	)`

	_, err := u.conn.Exec(
		ctx,query,
		user.UserID,
		user.Fullname,
		user.Username,
		user.Gmail,
	)
	if err != nil{
		log.Println("error on executing:",err)
	}
	return nil
}
func (u *UserRepo) GetUsers(ctx context.Context, limit, page string) ([]models.User, error) {
	return nil, nil
}
func (u *UserRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}
func (u *UserRepo) DeleteUserById(ctx context.Context, user models.User) error {
	return nil
}
