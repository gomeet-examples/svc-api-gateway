package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func TestCreateProfile(t *testing.T) {
	server := newApiGatewayServer()
	ctx := context.Background()

	req := &pb.ProfileCreationRequest{}
	// You can generate a fake request see https://github.com/gomeet/go-proto-gomeetfaker
	// req := &pb.ProfileCreationRequest{}
	res, err := server.CreateProfile(ctx, req)
	assert.Nil(t, err, "CreateProfile: error on call")
	assert.NotNil(t, res, "CreateProfile: error on call")

	// Do something useful tests with req and res
	// for example :
	// assert.Equal(t, req.GetUuid(), res.GetUuid(), "CreateProfile: Uuid field in response must be the same as that of the request")
}
