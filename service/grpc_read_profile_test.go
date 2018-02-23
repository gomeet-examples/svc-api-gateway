package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func TestReadProfile(t *testing.T) {
	server := newApiGatewayServer()
	ctx := context.Background()

	req := &pb.ProfileRequest{}
	// You can generate a fake request see https://github.com/gomeet/go-proto-gomeetfaker
	// req := &pb.ProfileRequest{}
	res, err := server.ReadProfile(ctx, req)
	assert.Nil(t, err, "ReadProfile: error on call")
	assert.NotNil(t, res, "ReadProfile: error on call")

	// Do something useful tests with req and res
	// for example :
	// assert.Equal(t, req.GetUuid(), res.GetUuid(), "ReadProfile: Uuid field in response must be the same as that of the request")
}
