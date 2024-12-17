package otp

import (
	desc "github.com/webbsalad/storya-otp-backend/internal/pb/github.com/webbsalad/storya-otp-backend/otp"
	service "github.com/webbsalad/storya-otp-backend/internal/service/otp"
)

type Implementation struct {
	desc.UnimplementedOtpServiceServer

	OtpService service.Service
}

func NewImplementation(otpService service.Service) desc.OtpServiceServer {
	return &Implementation{
		OtpService: otpService,
	}
}
