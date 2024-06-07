package db

import (
	"app/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

func ConnToDb(cfg config.Config) (*pgx.Conn, error){
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.PsqlConfig.User,
		cfg.PsqlConfig.Password,
		cfg.PsqlConfig.Host,
		cfg.PsqlConfig.Port,
		cfg.PsqlConfig.Name,
	)
	return pgx.Connect(ctx,dbUrl)
}

