package api

import (
	"github.com/brocaar/lora-app-server/internal/handler/httphandler"
	"github.com/brocaar/lora-app-server/internal/storage"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var errToCode = map[error]codes.Code{
	storage.ErrAlreadyExists:             codes.AlreadyExists,
	storage.ErrDoesNotExist:              codes.NotFound,
	storage.ErrApplicationInvalidName:    codes.InvalidArgument,
	storage.ErrNodeInvalidName:           codes.InvalidArgument,
	storage.ErrNodeMaxRXDelay:            codes.InvalidArgument,
	storage.ErrCFListTooManyChannels:     codes.InvalidArgument,
	storage.ErrUserInvalidUsername:       codes.InvalidArgument,
	storage.ErrUserPasswordLength:        codes.InvalidArgument,
	storage.ErrInvalidUsernameOrPassword: codes.Unauthenticated,
	httphandler.ErrInvalidHeaderName:     codes.InvalidArgument,
}

func errToRPCError(err error) error {
	cause := errors.Cause(err)
	code, ok := errToCode[cause]
	if !ok {
		code = codes.Unknown
	}
	return grpc.Errorf(code, cause.Error())
}
