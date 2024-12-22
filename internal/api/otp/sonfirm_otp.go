package otp

import (
	"context"

	"github.com/webbsalad/storya-otp-backend/internal/convertor"
	"github.com/webbsalad/storya-otp-backend/internal/pb/github.com/webbsalad/storya-otp-backend/otp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ConfirmOtp(ctx context.Context, req *otp.ConfirmOtpRequest) (*otp.ConfirmOtpResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	emailID, err := i.OtpService.ConfirmOtp(ctx, req.GetEmail(), req.GetOtpCode())
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &otp.ConfirmOtpResponse{
		EmailId: emailID.String(),
	}, nil
}
