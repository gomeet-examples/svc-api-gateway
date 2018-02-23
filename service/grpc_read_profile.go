package service

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gomeetContext "github.com/gomeet/gomeet/utils/context"
	"github.com/gomeet/gomeet/utils/log"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
	svcProfilePb "github.com/gomeet-examples/svc-profile/pb"
)

func castReadProfileRequest(req *pb.ProfileRequest) (*svcProfilePb.ProfileRequest, error) {
	if req == nil {
		return nil, errors.New("Invalid request (nil)")
	}

	return &svcProfilePb.ProfileRequest{Uuid: req.GetUuid()}, nil
}

func castReadProfileResponse(res *svcProfilePb.ProfileInfo) (*pb.ProfileInfo, error) {
	if res == nil {
		return nil, errors.New("Invalid response (nil)")
	}

	profileInfo := &pb.ProfileInfo{}
	if res != nil {
		var gender pb.Genders

		switch res.GetGender() {
		case svcProfilePb.Genders_MALE:
			gender = pb.Genders_MALE
		case svcProfilePb.Genders_FEMALE:
			gender = pb.Genders_FEMALE
		default:
			gender = pb.Genders_UNKNOW
		}

		profileInfo.Uuid = res.GetUuid()
		profileInfo.Gender = gender
		profileInfo.Email = res.GetEmail()
		profileInfo.Name = res.GetName()
		profileInfo.Birthday = res.GetBirthday()
		profileInfo.CreatedAt = res.GetCreatedAt()
		profileInfo.UpdatedAt = res.GetUpdatedAt()
		profileInfo.DeletedAt = res.GetDeletedAt()
	}

	return profileInfo, nil
}

func (s *apiGatewayServer) ReadProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileInfo, error) {
	log.Debug(ctx, "service call", log.Fields{"req": req})

	// validate request
	if err := req.Validate(); err != nil {
		log.Warn(ctx, "invalid request", err, log.Fields{
			"req": req,
		})

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//Initialize profileClient
	if err := s.initProfileClient(); err != nil {
		log.Error(ctx, "Failed to init svc-echo client", err, log.Fields{})

		return nil, status.Errorf(codes.Unavailable, "Internal service unavailable")
	}

	//Initialize the sub service context
	svcCtx := gomeetContext.NewSubServiceContext(ctx)

	// cast request
	profileReq, err := castReadProfileRequest(req)
	if err != nil {
		log.Error(ctx, "castReadProfileRequest Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	//call profile.Read grpc service
	profileRes, err := s.svcProfileGrpcClient.Read(svcCtx, profileReq)
	if err != nil {
		log.Error(ctx, "Failed to svcProfileGrpcClient.Read call", err, log.Fields{
			"profileReq": profileReq,
		})
		if eStatus, ok := status.FromError(err); ok {
			switch eStatus.Code() {
			case codes.InvalidArgument:
				return nil, status.Errorf(codes.InvalidArgument, eStatus.Message())
			}
		}

		return nil, status.Errorf(codes.Unavailable, "Internal service error")
	}

	res, err := castReadProfileResponse(profileRes)
	if err != nil {
		log.Error(ctx, "castReadProfileResponse Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	return res, nil
}
