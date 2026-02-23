// Copyright 2024 - Hanzo AI, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package cli

import (
	"os"
)

// Environment variables used by the KES CLI.
const (
	// EnvServer is the server endpoint the client uses. If not set,
	// clients will use '127.0.0.1:7373'.
	EnvServer = "S3_KES_SERVER"

	// EnvAPIKey is used by the client to authenticate to the server.
	EnvAPIKey = "S3_KES_API_KEY"

	// EnvPrivateKey is the path to the client private key file.
	EnvPrivateKey = "S3_KES_KEY_FILE"

	// EnvCertificate is the path to the client certificate file.
	EnvCertificate = "S3_KES_CERT_FILE"
)

// Env retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func Env(key string) string {
	switch key {
	default:
		return os.Getenv(key)
	case EnvServer:
		const (
			EnvServerLegacy = "KES_SERVER"
			EnvServerS3     = "S3_KMS_KES_ENDPOINT"
			DefaultServer   = "127.0.0.1:7373"
		)
		if s, ok := os.LookupEnv(EnvServer); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvServerLegacy); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvServerS3); ok {
			return s
		}
		return DefaultServer

	case EnvAPIKey:
		const (
			EnvAPIKeyLegacy = "KES_API_KEY"
			EnvAPIKeyS3     = "S3_KMS_KES_API_KEY"
		)
		if s, ok := os.LookupEnv(EnvAPIKey); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvAPIKeyLegacy); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvAPIKeyS3); ok {
			return s
		}
		return ""

	case EnvPrivateKey:
		const (
			EnvPrivateKeyLegacy = "KES_CLIENT_KEY"
			EnvPrivateKeyS3     = "S3_KES_CLIENT_KEY"
		)
		if s, ok := os.LookupEnv(EnvPrivateKey); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvPrivateKeyLegacy); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvPrivateKeyS3); ok {
			return s
		}
		return ""

	case EnvCertificate:
		const (
			EnvCertificateLegacy = "KES_CLIENT_CERT"
			EnvCertificateS3     = "S3_KES_CLIENT_CERT"
		)
		if s, ok := os.LookupEnv(EnvCertificate); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvCertificateLegacy); ok {
			return s
		}
		if s, ok := os.LookupEnv(EnvCertificateS3); ok {
			return s
		}
		return ""
	}
}
