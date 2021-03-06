// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/api-gateway.proto
// Package service provides gRPC/HTTP service registration
package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

var (
	name       = "svc-api-gateway"       // injected with -ldflags in Makefile
	version    = "latest"                // injected with -ldflags in Makefile
	controller *apiGatewayHTTPController // HTTP/1.1 controller
)

func init() {
	controller = newApiGatewayHTTPController()
}

// Service is the echo
type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewService return new svc-api-gateway service
func NewService() *Service {
	return &Service{
		Name:    name,
		Version: version,
	}
}

// RegisterGRPCServices register all grpc services
func (svc Service) RegisterGRPCServices(
	s *grpc.Server,
	jwtSecret string,
	caCert string,
	cert string,
	privKey string,
	// EXTRA : param
	// END EXTRA : param
	// SUB-SERVICES DEFINITION : param-address
	// svc{{SubServiceNamePascalCase}}Address string,
	svcEchoAddress string,
	svcProfileAddress string,
	// END SUB-SERVICES DEFINITION : param-address
) {
	log.WithFields(log.Fields{
		"jwtSecret": jwtSecret,
		"caCert":    caCert,
		"cert":      cert,
		"privKey":   privKey,
		// EXTRA : log
		// END EXTRA : log
		// SUB-SERVICES DEFINITION : log-address
		// "svc{{SubServiceNamePascalCase}}Address": svc{{SubServiceNamePascalCase}}Address,
		"svcEchoAddress":    svcEchoAddress,
		"svcProfileAddress": svcProfileAddress,
		// END SUB-SERVICES DEFINITION : log-address
	}).Debug("RegisterGRPCServices")
	pb.RegisterApiGatewayServer(s, &apiGatewayServer{
		jwtSecret:     jwtSecret,
		caCertificate: caCert,
		certificate:   cert,
		privateKey:    privKey,
		// EXTRA : register server
		// END EXTRA : register server
		// SUB-SERVICES DEFINITION : register-address-to-server
		// svc{{SubServiceNamePascalCase}}Address: svc{{SubServiceNamePascalCase}}Address,
		svcEchoAddress:    svcEchoAddress,
		svcProfileAddress: svcProfileAddress,
		// END SUB-SERVICES DEFINITION : register-address-to-server
	})
}

// RegisterHTTPServices register all http services
func (svc Service) RegisterHTTPServices(ctx context.Context, mux *mux.Router, addr string, opts []grpc.DialOption, jwtMiddleware *jwtmiddleware.JWTMiddleware) {
	// get server mux
	gwmux := runtime.NewServeMux()
	err := pb.RegisterApiGatewayHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("RegisterGRPCGateway error : %s\n", err)
	}

	// prometheus instrument handler
	instrf := prometheus.InstrumentHandlerFunc

	// HTTP/1.1 routes
	// status handler
	mux.
		HandleFunc("/status", instrf("Http.Status", controller.Status))

	mux.
		HandleFunc("/version", instrf("Http.Version", controller.Version))

	mux.
		HandleFunc("/", instrf("Http.Root", func(w http.ResponseWriter, r *http.Request) {
			// The "/" pattern matches everything not matched by previous handlers
			fmt.Fprintf(w, "%s-%s OK", svc.Name, svc.Version)
		}))

	// swagger doc handler
	mux.
		PathPrefix("/api/v1/swagger.json").
		Handler(instrf("Api.Swagger", controller.Swagger))

	// to declare an authenticated handler do something like this
	// if jwtMiddleware == nil {
	//   mux.
	//     PathPrefix("/<URL>").
	//     Handler(instrf("<METRICS_KEY>", controller.<HTTP_HANDLER>))
	// } else {
	//   mux.
	//     PathPrefix("/<URL>").
	//     Handler(negroni.New(
	//       negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
	//       negroni.Wrap(instrf("<METRICS_KEY>", controller.<HTTP_HANDLER>)),
	//     ))
	// }

	// it's not necessary to use secure middleware for gRPC calls
	// api gateway handlers with metrics instrumentations
	routeMap := map[string]string{
		"/api/v1/profile/read":    "Api.ReadProfile",
		"/api/v1/profile/list":    "Api.ListProfile",
		"/api/v1/profile/update":  "Api.UpdateProfile",
		"/api/v1/profile/delete":  "Api.DeleteProfile",
		"/api/v1/services/status": "Api.ServicesStatus",
		"/api/v1/profile/create":  "Api.CreateProfile",
		"/api/v1/version":         "Api.Version",
		"/api/v1/echo":            "Api.Echo",
	}
	for route, label := range routeMap {
		mux.PathPrefix(route).Handler(instrf(label, gwmux.ServeHTTP))
	}

	// prometheus metrics handler
	mux.
		Handle("/metrics", prometheus.Handler())
}
