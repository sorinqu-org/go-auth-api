package api

import (
	"github.com/jackc/pgx/v5"
)

type APIServer struct {
	addr string
	db   *pgx.Conn
}
