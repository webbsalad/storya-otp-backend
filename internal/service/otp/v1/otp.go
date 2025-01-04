package v1

import (
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/webbsalad/storya-otp-backend/internal/model"
)

func generateOtpKey(email string) (string, error) {
	otpKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "storya",
		AccountName: email,
		Digits:      otp.DigitsSix,
		Period:      60,
	})

	if err != nil {
		return "", fmt.Errorf("generate otp key: %w", err)
	}

	return otpKey.Secret(), nil
}

func generateOtpCode(otpKey string) (string, error) {
	otpCode, err := totp.GenerateCodeCustom(otpKey, time.Now(), totp.ValidateOpts{
		Period: 60,
		Digits: otp.DigitsSix,
	})
	if err != nil {
		return "", fmt.Errorf("generate otp code: %w", err)
	}

	return otpCode, nil
}

func validateOtp(otpKey, otpCode string) error {
	valid, err := totp.ValidateCustom(otpCode, otpKey, time.Now(), totp.ValidateOpts{
		Skew:   1,
		Period: 60,
		Digits: otp.DigitsSix,
	})
	if err != nil {
		return fmt.Errorf("validate otp code: %w", err)
	}

	if !valid {
		return model.ErrWrongOtp
	}

	return nil
}
