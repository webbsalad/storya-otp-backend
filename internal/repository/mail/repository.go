package mail

import "context"

type Repository interface {
	SendMail(ctx context.Context, email, otpCode string) error
}
