// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/api-gateway.proto
package cmd

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/gomeet/gomeet/utils/jwt"

	"github.com/gomeet-examples/svc-api-gateway/cmd/functest"
)

var (
	useEmbeddedServer bool
	useRandomPort     bool

	// funcTestCmd represents the functest command
	funcTestCmd = &cobra.Command{
		Use:   "functest",
		Short: "Runs functional tests on the service",
		Run: func(cmd *cobra.Command, args []string) {
			runFunctionalTests()
		},
	}
)

func init() {
	RootCmd.AddCommand(funcTestCmd)

	// force debug mode
	funcTestCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Force debug mode")

	// address flag (to serve all protocols on a single port)
	funcTestCmd.PersistentFlags().StringVarP(&serverAddress, "address", "a", "localhost:13000", "Multiplexed gRPC/HTTP server address")

	// gRPC address flag (to serve gRPC on a separate address)
	funcTestCmd.PersistentFlags().StringVar(&grpcServerAddress, "grpc-address", "", "gRPC server address")

	// HTTP/1.1 address flag (to serve HTTP on a separate address)
	funcTestCmd.PersistentFlags().StringVar(&httpServerAddress, "http-address", "", "HTTP server address")

	// CA certificate
	funcTestCmd.PersistentFlags().StringVar(&caCertificate, "ca", "", "X.509 certificate of the Certificate Authority (required for gRPC TLS support)")

	// gRPC client certificate
	funcTestCmd.PersistentFlags().StringVar(&serverCertificate, "cert", "", "X.509 certificate (required for gRPC TLS support)")

	// gRPC client private key
	funcTestCmd.PersistentFlags().StringVar(&serverPrivateKey, "key", "", "RSA private key (required for gRPC TLS support)")

	// JSON Web Token
	funcTestCmd.PersistentFlags().StringVar(&jwtToken, "jwt", "", "JSON Web Token (external server only)")

	// JWT secret signing key
	funcTestCmd.PersistentFlags().StringVar(&jwtSecret, "jwt-secret", "", "JSON Web Token secret signing key (embedded server only)")

	// request timeout
	funcTestCmd.PersistentFlags().IntVar(&timeoutSeconds, "timeout", 5, "Request timeout in seconds")

	// embedded server flag
	funcTestCmd.PersistentFlags().BoolVarP(&useEmbeddedServer, "embed-server", "e", false, "Embed the server to test")

	// random port server flag
	funcTestCmd.PersistentFlags().BoolVar(&useRandomPort, "random-port", false, "Use a random port for the embedded server")

}

// getFreePort asks the kernel for a free open port that is ready to use.
func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func runFunctionalTests() {
	if useEmbeddedServer {
		if debugMode {
			log.SetLevel(log.DebugLevel)
		} else {
			// by default for embedded server only panic are logged
			log.SetLevel(log.PanicLevel)
		}
		if useRandomPort {
			freePort, err := getFreePort()
			if err == nil {
				serverAddress = fmt.Sprintf("localhost:%d", freePort)
				grpcServerAddress = ""
				httpServerAddress = ""
			}
		}
		if grpcServerAddress != "" && httpServerAddress != "" {
			go serveOnMultipleAddresses(grpcServerAddress, httpServerAddress)
		} else {
			go serveOnSingleAddress(serverAddress)
		}
	}

	// if the embedded server is JWT-enabled, test clients require a valid token
	if useEmbeddedServer && jwtSecret != "" {
		var err error
		jwtToken, err = jwt.Create(
			"github.com/gomeet/gomeet",
			secretSigningKey,
			tokenLifetimeHours,
			subjectID,
			jwt.Claims{},
		)

		if err != nil {
			fmt.Printf("failed to create JWT : %v\n", err)
			os.Exit(1)
		}
	}

	testConfig := functest.FunctionalTestConfig{
		ServerAddress:     serverAddress,
		GrpcServerAddress: grpcServerAddress,
		HttpServerAddress: httpServerAddress,
		CaCertificate:     caCertificate,
		ClientCertificate: serverCertificate,
		ClientPrivateKey:  serverPrivateKey,
		TimeoutSeconds:    timeoutSeconds,
		JsonWebToken:      jwtToken,
	}

	failures := runFunctionalTestSession(testConfig)

	if len(failures) == 0 {
		fmt.Println("PASS")
		fmt.Println("ok\tfunctest is ok")

		os.Exit(0)
	} else {
		fmt.Printf("Test failures:\n")
		for idx, failure := range failures {
			fmt.Printf("%d) %s: %s\n", idx+1, failure.Procedure, failure.Message)
		}

		os.Exit(1)
	}
}

func appendFailures(acc []functest.TestFailure, newSlice []functest.TestFailure) []functest.TestFailure {
	for _, failure := range newSlice {
		acc = append(acc, failure)
	}

	return acc
}

func runFunctionalTestSession(config functest.FunctionalTestConfig) []functest.TestFailure {
	var failures []functest.TestFailure

	// gRPC services test
	failures = appendFailures(failures, functest.TestVersion(config))
	failures = appendFailures(failures, functest.TestHttpVersion(config))
	failures = appendFailures(failures, functest.TestServicesStatus(config))
	failures = appendFailures(failures, functest.TestHttpServicesStatus(config))
	failures = appendFailures(failures, functest.TestEcho(config))
	failures = appendFailures(failures, functest.TestHttpEcho(config))
	failures = appendFailures(failures, functest.TestCreateProfile(config))
	failures = appendFailures(failures, functest.TestHttpCreateProfile(config))
	failures = appendFailures(failures, functest.TestReadProfile(config))
	failures = appendFailures(failures, functest.TestHttpReadProfile(config))
	failures = appendFailures(failures, functest.TestListProfile(config))
	failures = appendFailures(failures, functest.TestHttpListProfile(config))
	failures = appendFailures(failures, functest.TestUpdateProfile(config))
	failures = appendFailures(failures, functest.TestHttpUpdateProfile(config))
	failures = appendFailures(failures, functest.TestDeleteProfile(config))
	failures = appendFailures(failures, functest.TestHttpDeleteProfile(config))
	// Extra http handler
	failures = appendFailures(failures, functest.TestHttpStatus(config))
	failures = appendFailures(failures, functest.TestHttpMetrics(config))
	failures = appendFailures(failures, functest.TestHttpSwagger(config))

	return failures
}
