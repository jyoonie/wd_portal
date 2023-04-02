package service

import (
	"database/sql"

	"go.uber.org/zap"
)

type Service struct {
	l  zap.Logger
	db *sql.DB
}
