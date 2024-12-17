package main

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
		return "", fmt.Errorf("generate otps: %w", err)
	}

	return otpKey.Secret(), nil
}

func generateOtpCode(otpKey string) (string, error) {
	otpCode, err := totp.GenerateCode(otpKey, time.Now())
	if err != nil {
		return "", fmt.Errorf("generate otp code: %w", err)
	}

	return otpCode, nil
}

func main() {
	fmt.Println(generateOtpKey("1@ex.com"))
	fmt.Println(generateOtpCode("BCYTJYTK6GDEF3MB543SRVJWIHOP3I6N"))
}
