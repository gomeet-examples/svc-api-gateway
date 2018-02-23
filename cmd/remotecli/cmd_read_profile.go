package remotecli

import (
	"errors"
	"fmt"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func (c *remoteCli) cmdReadProfile(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("Bad arguments : read_profile <uuid [string]>")
	}

	// request message
	var req *pb.ProfileRequest

	// decl req for no nil panic
	req = &pb.ProfileRequest{}

	// cast args[0] in req.Uuid - type TYPE_STRING to go type string
	req.Uuid = args[0]

	// message validation - github.com/mwitkow/go-proto-validators
	if reqValidator, ok := interface{}(*req).(interface {
		Validate() error
	}); ok {
		if err := reqValidator.Validate(); err != nil {
			return "", err
		}
	}

	// sending message to server
	r, err := c.c.ReadProfile(c.ctx, req)
	if err != nil {
		return "", fmt.Errorf("ReadProfile service call fail - %v", err)
	}

	return fmt.Sprintf("ReadProfile: %v", r), nil
}