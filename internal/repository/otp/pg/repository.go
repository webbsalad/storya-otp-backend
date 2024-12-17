package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-otp-backend/internal/repository/otp"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (otp.Repository, error) {
	return &Repository{db: db}, nil
}
