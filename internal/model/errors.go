package model

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist     = errors.New("already exist")
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotFound         = errors.New("not found")
)

var (
	ErrOtpKeyNotFound = fmt.Errorf("user not found: %w", ErrNotFound)

	ErrWrongOtp = fmt.Errorf("wrong otp code: %w", ErrPermissionDenied)

	ErrOtpKeyAlreadyExist = fmt.Errorf("otp key already exist: %w", ErrAlreadyExist)
)
