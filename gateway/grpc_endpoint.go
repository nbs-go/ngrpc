package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EndpointRegisterHandlerFunc = func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type GrpcEndpoint struct {
	Name            string
	RegisterHandler EndpointRegisterHandlerFunc
	Hostname        string
	Insecure        bool
	SSLCertificate  string
}

func RegisterEndpoints(ctx context.Context, gw *runtime.ServeMux, endpoints []GrpcEndpoint) error {
	for _, endpoint := range endpoints {
		if !endpoint.Insecure {
			return fmt.Errorf("ngrpc/gateway: secure grpc server is not yet supported. EndpointName=%s", endpoint.Name)
		}

		hOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		// Register endpoints
		err := endpoint.RegisterHandler(ctx, gw, endpoint.Hostname, hOpts)
		if err != nil {
			return err
		}
		log.Tracef("gRPC Endpoint registered. EndpointName=%s, Hostname=%s", endpoint.Name, endpoint.Hostname)
	}
	return nil
}
