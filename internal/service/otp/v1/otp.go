package v1

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func generateOtpKey(email string) (string, error) {
	otpKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "storya",
		AccountName: email,
	})

	if err != nil {
		return "", fmt.Errorf("generate otp key: %w", err)
	}

	return string(otpKey.Secret()), nil
}

func generateOtpCode(otpKey string) (string, error) {
	otpCode, err := totp.GenerateCode(otpKey, time.Now())
	if err != nil {
		return "", fmt.Errorf("generate otp code: %w", err)
	}

	return otpCode, nil
}
