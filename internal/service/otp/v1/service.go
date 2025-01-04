package v1

import (
	"github.com/webbsalad/storya-otp-backend/internal/config"
	"github.com/webbsalad/storya-otp-backend/internal/repository/mail"
	"github.com/webbsalad/storya-otp-backend/internal/repository/otp"
	otp_service "github.com/webbsalad/storya-otp-backend/internal/service/otp"
)

type Service struct {
	otpRepository  otp.Repository
	mailRepository mail.Repository
	config         config.Config
}

func NewService(otpRepository otp.Repository, mailRepository mail.Repository, config config.Config) otp_service.Service {
	return &Service{
		otpRepository:  otpRepository,
		mailRepository: mailRepository,
		config:         config,
	}
}
