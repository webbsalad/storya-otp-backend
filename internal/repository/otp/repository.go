package otp

import (
	"context"

	"github.com/webbsalad/storya-otp-backend/internal/model"
)

type Repository interface {
	GetOtpKey(ctx context.Context, email string) (string, error)
	StoreOtpKey(ctx context.Context, email, otpKey string) (string, error)
	ConfirmEmail(ctx context.Context, email string) (model.EmailID, error)
}
