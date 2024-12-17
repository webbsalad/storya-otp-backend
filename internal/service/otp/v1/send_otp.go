package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/webbsalad/storya-otp-backend/internal/model"
)

func (s *Service) SendOtp(ctx context.Context, email string) error {
	otpKey, err := s.otpRepository.GetOtpKey(ctx, email)
	if err != nil {
		if errors.Is(err, model.ErrOtpKeyNotFound) {
			otpKey, err = generateOtpKey(email)
			if err != nil {
				return fmt.Errorf("otp key: %w", err)
			}
		}
	}

	otpKey, err = s.otpRepository.StoreOtpKey(ctx, email, string(otpKey))
	if err != nil {
		return fmt.Errorf("store otp code: %w", err)
	}

	otpCode, err := generateOtpCode(otpKey)
	if err != nil {
		return fmt.Errorf("otp code: %w", err)
	}

	if err = sendMail(email, otpCode); err != nil {
		return fmt.Errorf("sending email: %w", err)
	}

	return nil

}
