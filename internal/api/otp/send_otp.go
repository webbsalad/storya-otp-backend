package otp

import (
	"context"

	"github.com/webbsalad/storya-otp-backend/internal/pb/github.com/webbsalad/storya-otp-backend/otp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendOtp(ctx context.Context, req *otp.SendOtpRequest) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	if err := i.OtpService.SendOtp(ctx, req.GetEmail()); err != nil {
		return nil, status.Errorf(codes.Internal, "server error: %v", err)
	}

	return &emptypb.Empty{}, nil
}
