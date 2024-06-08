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
		INSERT INTO users(
			user_id,
			fullname, 
			username, 
			gmail,
			password 
		) VALUES (
			$1, $2, $3, $4, $5
		)`

	_, err := u.conn.Exec(
		ctx, query,
		user.UserID,
		user.Fullname,
		user.Username,
		user.Gmail,
		user.Password,
	)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	return nil
}

// func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error {
// 	query := `
// 		INSERT INTO users(
// 		user_id,
// 		fullname, 
// 		username, 
// 		gmail,
// 		password 
// 	)VALUES (
// 		$1 ,$2, $3, $4, $5
// 	)`

// 	_, err := u.conn.Exec(
// 		ctx,query,
// 		user.UserID,
// 		user.Fullname,
// 		user.Username,
// 		user.Gmail,
// 		user.Password,
// 	)
// 	if err != nil{
// 		log.Println("error on executing:",err)
// 		return err
// 	}
// 	return nil
// }

func (u *UserRepo) GetUsers(ctx context.Context, limit, page string) ([]models.User, error) {
	var users []models.User

	query := `
		SELECT user_id, fullname, username, gmail, password
		FROM users
		LIMIT $1 OFFSET $2
	`

	rows, err := u.conn.Query(ctx, query, limit, page)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Fullname, &user.Username, &user.Gmail, &user.Password); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error retrieving rows:", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	var user models.User

	query := `
		SELECT user_id, fullname, username, gmail, password
		FROM users
		WHERE user_id = $1
		LIMIT 1
	`

	err := u.conn.QueryRow(ctx, query, userId).Scan(&user.UserID, &user.Fullname, &user.Username, &user.Gmail, &user.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user models.User) error {
	query := `
		UPDATE users
		SET fullname = $1, username = $2, gmail = $3, password = $4
		WHERE user_id = $5
	`

	_, err := u.conn.Exec(ctx, query, user.Fullname, user.Username, user.Gmail, user.Password, user.UserID)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (u *UserRepo) DeleteUserById(ctx context.Context, userId models.User) error {
	query := `
		DELETE FROM users
		WHERE user_id = $1
	`

	_, err := u.conn.Exec(ctx, query, userId)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}
