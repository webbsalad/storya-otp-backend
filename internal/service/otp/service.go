package service

import "context"

type Service interface {
	SendOtp(ctx context.Context, email string) error
}
