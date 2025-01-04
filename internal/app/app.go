package app

import (
	"github.com/webbsalad/storya-otp-backend/internal/api/otp"
	"github.com/webbsalad/storya-otp-backend/internal/config"
	pb "github.com/webbsalad/storya-otp-backend/internal/pb/github.com/webbsalad/storya-otp-backend/otp"
	mail_repository "github.com/webbsalad/storya-otp-backend/internal/repository/mail/api"
	otp_repository "github.com/webbsalad/storya-otp-backend/internal/repository/otp/pg"

	v1 "github.com/webbsalad/storya-otp-backend/internal/service/otp/v1"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			initDB,
		),
		grpcOptions(),
		serviceOptions(),
	)
}

func serviceOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			otp.NewImplementation,
			otp_repository.NewRepository,
			mail_repository.NewRepository,
			v1.NewService,
		),
		fx.Invoke(
			pb.RegisterOtpServiceServer,
		),
	)
}
