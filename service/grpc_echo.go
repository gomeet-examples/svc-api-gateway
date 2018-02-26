package service

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gomeetContext "github.com/gomeet/gomeet/utils/context"
	"github.com/gomeet/gomeet/utils/log"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
	svcEchoPb "github.com/gomeet-examples/svc-echo/pb"
)

func castEchoRequest(req *pb.EchoRequest) (*svcEchoPb.EchoRequest, error) {
	if req == nil {
		return nil, errors.New("Invalid request (nil)")
	}

	return &svcEchoPb.EchoRequest{Uuid: req.GetUuid(), Content: req.GetContent()}, nil
}

func castEchoResponse(res *svcEchoPb.EchoResponse) (*pb.EchoResponse, error) {
	if res == nil {
		return nil, errors.New("Invalid response (nil)")
	}

	return &pb.EchoResponse{Uuid: res.GetUuid(), Content: res.GetContent()}, nil
}

func (s *apiGatewayServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Debug(ctx, "service call", log.Fields{"req": req})

	// validate request
	if err := req.Validate(); err != nil {
		log.Warn(ctx, "invalid request", err, log.Fields{
			"req": req,
		})

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//Initialize echoClient
	if err := s.initEchoClient(); err != nil {
		log.Error(ctx, "Failed to init svc-echo client", err, log.Fields{})

		return nil, status.Errorf(codes.Unavailable, "Internal service unavailable")
	}

	//Initialize the sub service context
	svcCtx := gomeetContext.NewSubServiceContext(ctx)

	// cast request
	echoReq, err := castEchoRequest(req)
	if err != nil {
		log.Error(ctx, "castEchoRequest Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	//call echo.Echo grpc service
	echoRes, err := s.svcEchoGrpcClient.Echo(svcCtx, echoReq)
	if err != nil {
		log.Error(ctx, "Failed to svcEchoGrpcClient.Echo call", err, log.Fields{
			"echoReq": echoReq,
		})
		if eStatus, ok := status.FromError(err); ok {
			switch eStatus.Code() {
			case codes.InvalidArgument:
				return nil, status.Errorf(codes.InvalidArgument, eStatus.Message())
			}
		}

		return nil, status.Errorf(codes.Unavailable, "Internal service error")
	}

	res, err := castEchoResponse(echoRes)
	if err != nil {
		log.Error(ctx, "castEchoResponse Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	return res, nil
}
