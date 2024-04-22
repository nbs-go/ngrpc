package ngrpc

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"google.golang.org/grpc/credentials"
)

// X509KeyPairFromBase64 parses a public/private key pair from a pair of
// PEM encoded data. On successful return, Certificate. Leaf will be nil because
// the parsed form of the certificate is not retained.
//
// Encoded certificate and key is base64 encoded string
func X509KeyPairFromBase64(certB64, keyB64 string) (*tls.Certificate, error) {
	// Decode certificate
	cb, err := base64.StdEncoding.DecodeString(certB64)
	if err != nil {
		return nil, err
	}

	// Decode private key
	kb, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return nil, err
	}

	// Load key pair
	cert, err := tls.X509KeyPair(cb, kb)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// NewClientTLSFromBase64 constructs TLS credentials from the provided root
// certificate authority certificate file(s) to validate server connections. If
// certificates to establish the identity of the client need to be included in
// the credentials (eg: for mTLS), use NewTLS instead, where a complete
// tls.Config can be specified.
// serverNameOverride is for testing only. If set to a non-empty string,
// it will override the virtual host name of authority (e.g. :authority header
// field) in requests.
func NewClientTLSFromBase64(certB64, serverNameOverride string) (credentials.TransportCredentials, error) {
	// Decode certificate
	cb, err := base64.StdEncoding.DecodeString(certB64)
	if err != nil {
		return nil, err
	}

	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(cb) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}
	return credentials.NewTLS(&tls.Config{ServerName: serverNameOverride, RootCAs: cp, MinVersion: tls.VersionTLS12}), nil
}
