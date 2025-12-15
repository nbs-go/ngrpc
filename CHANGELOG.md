# CHANGELOG

## v0.2.0

- fix: Upgrade go to v1.24, upgrade all dependencies
- fix(proto): Generate proto using protoc-gen-go v1.36.10 and protoc v6.31.1

## v0.1.1

- fix: Upgrade Go dependency version to 1.19

## v0.1.0

- feat(gateway): Add gRPC Gateway Handlers
- feat(unary): Add RequestInterceptor for unary grpc server
- feat(unary): Add RecoveryInterceptor for unary grpc server
- feat: Add common grpc to http error mapping
- feat: Add RequestMetadata
- feat: Add RequestHandler interface
- feat: Add Listen Port composer
- feat: Add gRPC Header value getter
- feat: Add ErrorDetails for gRPC Error
- feat: Add certificate loader from encoded base64

