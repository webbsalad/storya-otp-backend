package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-otp-backend/internal/model"
)

func (s *Service) ConfirmOtp(ctx context.Context, email, otp string) (model.EmailID, error) {
	storedKey, err := s.otpRepository.GetOtpKey(ctx, email)
	if err != nil {
		return model.EmailID{}, fmt.Errorf("get otp: %w", err)
	}

	if err := validateOtp(storedKey, otp); err != nil {
		return model.EmailID{}, fmt.Errorf("code comparison: %w", err)
	}

	emailID, err := s.otpRepository.ConfirmEmail(ctx, email)
	if err != nil {
		return model.EmailID{}, fmt.Errorf("confirm email: %w", err)
	}

	return emailID, nil
}
