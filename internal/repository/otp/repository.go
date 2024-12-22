package otp

import "context"

type Repository interface {
	GetOtpKey(ctx context.Context, email string) (string, error)
	StoreOtpKey(ctx context.Context, email, otpKey string) (string, error)
}
