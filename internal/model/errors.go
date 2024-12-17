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

	ErrOtpKeyAlreadyExist = fmt.Errorf("otp key already exist: %w", ErrAlreadyExist)
)
