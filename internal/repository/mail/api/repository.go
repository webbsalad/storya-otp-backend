package api

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-otp-backend/internal/repository/mail"
)

type Repository struct {
}

func NewRepository() (mail.Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) SendMail(ctx context.Context, email, otpCode string) error {
	// TODO реализовать логику
	_ = email
	fmt.Println(otpCode)
	return nil
}
