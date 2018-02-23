package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func TestListProfile(t *testing.T) {
	server := newApiGatewayServer()
	ctx := context.Background()

	req := &pb.ProfileListRequest{}
	// You can generate a fake request see https://github.com/gomeet/go-proto-gomeetfaker
	// req := &pb.ProfileListRequest{}
	res, err := server.ListProfile(ctx, req)
	assert.Nil(t, err, "ListProfile: error on call")
	assert.NotNil(t, res, "ListProfile: error on call")

	// Do something useful tests with req and res
	// for example :
	// assert.Equal(t, req.GetUuid(), res.GetUuid(), "ListProfile: Uuid field in response must be the same as that of the request")
}
