package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-otp-backend/internal/model"
	"github.com/webbsalad/storya-otp-backend/internal/repository/otp"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (otp.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) GetOtpKey(ctx context.Context, email string) (string, error) {
	var otpKey string

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Select("otp_key").
		From("otp_keys").
		Where(
			sq.Eq{"email": email},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("build query: %w", err)
	}

	if err = r.db.GetContext(ctx, &otpKey, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrOtpKeyNotFound
		}

		return "", fmt.Errorf("get otp key: %w", err)
	}

	return otpKey, nil

}

func (r *Repository) StoreOtpKey(ctx context.Context, email, otpKey string) (string, error) {
	var storedKey string

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("otp_keys").
		Columns("email", "otp_key").
		Values(email, otpKey).
		Suffix("ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email RETURNING id")

	q, args, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("build query: %w", err)
	}

	if err = r.db.QueryRowContext(ctx, q, args...).Scan(&storedKey); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			otpKey = storedKey
		} else {
			return "", fmt.Errorf("execute query: %w", err)
		}

	}

	return otpKey, nil
}

func (r *Repository) ConfirmEmail(ctx context.Context, email string) (model.EmailID, error) {
	var storedEmailID string

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("confirmed_emails").
		Columns("email").
		Values(email).
		Suffix("ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email RETURNING id")

	q, args, err := query.ToSql()
	if err != nil {
		return model.EmailID{}, fmt.Errorf("build query: %w", err)
	}

	if err = r.db.QueryRowContext(ctx, q, args...).Scan(&storedEmailID); err != nil {
		return model.EmailID{}, fmt.Errorf("execute query: %w", err)
	}

	emailID, err := model.EmailIDFromString(storedEmailID)
	if err != nil {
		return model.EmailID{}, fmt.Errorf("convert str to email id: %w", err)
	}

	return emailID, nil

}
