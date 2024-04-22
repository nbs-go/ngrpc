package ngrpc

import (
	"github.com/nbs-go/errx"
	"google.golang.org/grpc/codes"
)

func WithStatus(status codes.Code) errx.SetOptionFn {
	return errx.AddMetadata(KeyGrpcStatus, status)
}

func OverrideMessage(message string) errx.SetOptionFn {
	return errx.AddMetadata(KeyOverrideMessage, message)
}

// Define Common Errors

var b = errx.NewBuilder(pkgNamespace)

var InternalError = b.NewError("500", "Internal Error",
	WithStatus(codes.Internal),
)

var BadRequestError = b.NewError("400", "Bad Request",
	WithStatus(codes.InvalidArgument),
)

var UnauthorizedError = b.NewError("401", "Unauthorized",
	WithStatus(codes.Unauthenticated),
)

var ForbiddenError = b.NewError("403", "Forbidden",
	WithStatus(codes.PermissionDenied),
)

var NotFoundError = b.NewError("404", "Not Found",
	WithStatus(codes.NotFound),
)

var CancelError = b.NewError("408", "Request Canceled",
	WithStatus(codes.Canceled),
)
