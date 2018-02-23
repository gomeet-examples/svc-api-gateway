package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func TestUpdateProfile(t *testing.T) {
	server := newApiGatewayServer()
	ctx := context.Background()

	req := &pb.ProfileInfo{}
	// You can generate a fake request see https://github.com/gomeet/go-proto-gomeetfaker
	// req := &pb.ProfileInfo{}
	res, err := server.UpdateProfile(ctx, req)
	assert.Nil(t, err, "UpdateProfile: error on call")
	assert.NotNil(t, res, "UpdateProfile: error on call")

	// Do something useful tests with req and res
	// for example :
	// assert.Equal(t, req.GetUuid(), res.GetUuid(), "UpdateProfile: Uuid field in response must be the same as that of the request")
}
