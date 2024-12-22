package service

import (
	"context"

	"github.com/webbsalad/storya-otp-backend/internal/model"
)

type Service interface {
	SendOtp(ctx context.Context, email string) error
	ConfirmOtp(ctx context.Context, email, otp string) (model.EmailID, error)
}
