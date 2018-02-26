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

func castListProfileRequest(req *pb.ProfileListRequest) (*svcProfilePb.ProfileListRequest, error) {
	if req == nil {
		return nil, errors.New("Invalid request (nil)")
	}

	var gender svcProfilePb.Genders

	switch req.GetGender() {
	case pb.Genders_MALE:
		gender = svcProfilePb.Genders_MALE
	case pb.Genders_FEMALE:
		gender = svcProfilePb.Genders_FEMALE
	default:
		gender = svcProfilePb.Genders_UNKNOW
	}

	return &svcProfilePb.ProfileListRequest{
		PageNumber:         req.GetPageNumber(),
		PageSize:           req.GetPageSize(),
		Order:              req.GetOrder(),
		ExcludeSoftDeleted: req.GetExcludeSoftDeleted(),
		SoftDeletedOnly:    req.GetSoftDeletedOnly(),
		Gender:             gender,
	}, nil
}

func castListProfileResponse(res *svcProfilePb.ProfileList) (*pb.ProfileList, error) {
	if res == nil {
		return nil, errors.New("Invalid response (nil)")
	}

	var profiles []*pb.ProfileInfo

	for _, profileProfileInfo := range res.GetProfiles() {
		var gender pb.Genders

		switch profileProfileInfo.GetGender() {
		case svcProfilePb.Genders_MALE:
			gender = pb.Genders_MALE
		case svcProfilePb.Genders_FEMALE:
			gender = pb.Genders_FEMALE
		default:
			gender = pb.Genders_UNKNOW
		}

		profiles = append(
			profiles,
			&pb.ProfileInfo{
				Uuid:      profileProfileInfo.GetUuid(),
				Gender:    gender,
				Email:     profileProfileInfo.GetEmail(),
				Name:      profileProfileInfo.GetName(),
				Birthday:  profileProfileInfo.GetBirthday(),
				CreatedAt: profileProfileInfo.GetCreatedAt(),
				UpdatedAt: profileProfileInfo.GetUpdatedAt(),
				DeletedAt: profileProfileInfo.GetDeletedAt(),
			},
		)
	}

	return &pb.ProfileList{
		ResultSetSize: res.GetResultSetSize(),
		HasMore:       res.GetHasMore(),
		Profiles:      profiles,
	}, nil
}

func (s *apiGatewayServer) ListProfile(ctx context.Context, req *pb.ProfileListRequest) (*pb.ProfileList, error) {
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
	profileReq, err := castListProfileRequest(req)
	if err != nil {
		log.Error(ctx, "castListProfileRequest Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	//call profile.List grpc service
	profileRes, err := s.svcProfileGrpcClient.List(svcCtx, profileReq)
	if err != nil {
		log.Error(ctx, "Failed to svcProfileGrpcClient.List call", err, log.Fields{
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

	res, err := castListProfileResponse(profileRes)
	if err != nil {
		log.Error(ctx, "castListProfileResponse Failed", err, log.Fields{
			"req": req,
		})

		return nil, status.Errorf(codes.Internal, "Internal service error")
	}

	return res, nil
}
