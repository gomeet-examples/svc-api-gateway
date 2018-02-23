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

func castDeleteProfileRequest(req *pb.ProfileRequest) (*svcProfilePb.ProfileRequest, error) {
	if req == nil {
		return nil, errors.New("Invalid request (nil)")
	}

	return &svcProfilePb.ProfileRequest{Uuid: req.GetUuid()}, nil
}

func castDeleteProfileResponse(res *svcProfilePb.ProfileResponse) (*pb.ProfileResponse, error) {
	if res == nil {
		return nil, errors.New("Invalid response (nil)")
	}

	profileInfo := &pb.ProfileInfo{}
	profileProfileInfo := res.GetInfo()

	if profileProfileInfo != nil {
		var gender pb.Genders

		switch profileProfileInfo.GetGender() {
		case svcProfilePb.Genders_MALE:
			gender = pb.Genders_MALE
		case svcProfilePb.Genders_FEMALE:
			gender = pb.Genders_FEMALE
		default:
			gender = pb.Genders_UNKNOW
		}

		profileInfo.Uuid = profileProfileInfo.GetUuid()
		profileInfo.Gender = gender
		profileInfo.Email = profileProfileInfo.GetEmail()
		profileInfo.Name = profileProfileInfo.GetName()
		profileInfo.Birthday = profileProfileInfo.GetBirthday()
		profileInfo.CreatedAt = profileProfileInfo.GetCreatedAt()
		profileInfo.UpdatedAt = profileProfileInfo.GetUpdatedAt()
		profileInfo.DeletedAt = profileProfileInfo.GetDeletedAt()
	}

	return &pb.ProfileResponse{Ok: res.GetOk(), Info: profileInfo}, nil
}

func (s *apiGatewayServer) DeleteProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
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
	profileReq, err := castDeleteProfileRequest(req)
	if err != nil {
		log.Error(ctx, "castDeleteProfileRequest Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	//call profile.Delete grpc service
	profileRes, err := s.svcProfileGrpcClient.SoftDelete(svcCtx, profileReq)
	if err != nil {
		log.Error(ctx, "Failed to svcProfileGrpcClient.SoftDelete call", err, log.Fields{
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

	res, err := castDeleteProfileResponse(profileRes)
	if err != nil {
		log.Error(ctx, "castDeleteProfileResponse Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	return res, nil
}
